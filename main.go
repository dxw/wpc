package main

import (
    _ "fmt"
    "os"
    "io/ioutil"
    "bytes"
    "flag"
)


func main() {
    /* Go won't allow the same arg/flag construction we had before, so all flags instead */
    project := flag.String("project", "", "Project name")
    multisite := flag.Bool("multisite", false, "Make it a multisite")
    flag.Parse()
    if *project == "" {
        println("You must set a value for the project argument")
        os.Exit(1)
    }
    if *multisite == true {
        println("MULTISITE")
    }

    makedir("bin")
    makedir("setup")
    makedir("setup/content")

    createDockerCompose(*project)
    createBinWp()
    createBinConsole()
    createBinSetup()
    createExternal()
    createInternal(*multisite)
}

func createDockerCompose(project string) {
    dockerComposeContents := []byte(
`version: "3"

volumes:
  mysql_data_!!!PROJECTNAME!!!:

services:
  mailcatcher:
    image: schickling/mailcatcher
    ports:
      - "1080:1080"

  beanstalk:
    image: schickling/beanstalkd
    ports:
      - "11300:11300"

  beanstalkd_console:
    image: agaveapi/beanstalkd-console
    ports:
      - "2080:80"
    environment:
      BEANSTALKD_HOST: beanstalk
      BEANSTALKD_PORT: 11300

  mysql:
    image: mysql
    ports:
      - "3306:3306"
    volumes:
      - mysql_data_!!!PROJECTNAME!!!:/var/lib/mysql
    environment:
      MYSQL_DATABASE: wordpress
      MYSQL_ROOT_PASSWORD: foobar

  wordpress:
    image: thedxw/wpc-wordpress
    ports:
      - "80:80"
    links:
      - mysql
      - mailcatcher
      - beanstalk
    volumes:
      - .:/usr/src/app
      - ./wp-content:/var/www/html/wp-content`)
    dockerComposeContents = findAndReplace(dockerComposeContents, []byte("!!!PROJECTNAME!!!"), []byte(project))
    creatingNoexec("world/docker-compose.yml", dockerComposeContents)
}

func createBinWp() {
    contents := []byte(
`#!/bin/sh
set -e

FLAGS=

# Add -t flag iff STDIN is a TTY
if [ -t 0 ]; then
  FLAGS=-t
fi

CONTAINER=` + "`" + `docker-compose ps -q wordpress` + "`" + `

# We can't use docker-compose here because docker-compose exec is equivalent
# to docker exec -ti and docker-compose exec -T is equivalent to
# docker exec. There is no docker-compose equivalent to docker exec -i.
#
# Issue: https://github.com/docker/compose/issues/3352

docker exec -i ${FLAGS} ${CONTAINER} wp "${@}"`)
    creating("bin/wp", contents)
}

func createBinConsole() {
    contents := []byte(
`#!/bin/sh
set -e

exec docker-compose exec wordpress bash`)
    creating("bin/console", contents)
}

func createBinSetup() {
    contents := []byte(
`#!/bin/sh
set -e
#
# Runs all site setup scripts

` + "`" + `dirname $0` + "`" + `/../setup/external.sh
docker-compose exec wordpress /usr/src/app/setup/internal.sh`)
    creating("bin/setup", contents)
}

func createExternal() {
    contents := []byte(
`#!/bin/sh
set -e

##
## This code will be run during setup, OUTSIDE the container.
##
## Because ` + "`" + `whippet` + "`" + ` running inside the container wouldn't be able to connect
## to private repositories.
##

if test -f whippet.json; then
  whippet deps install
fi`)
    creating("setup/external.sh", contents)
}

func createInternal(multisite bool) {
    contents := []byte(
`#!/bin/sh
set -e

##
## This code will be run during setup, INSIDE the container.
##

##############
# Config
##############
title="Your site title here"
theme=your-theme-slug
plugins="a-space-separated list-of plugins-to-activate"
content=/usr/src/app/setup/content

wp core !!!INSTALLTYPE!!! --skip-email --admin_user=admin --admin_password=admin --admin_email=admin@localhost.invalid --url=http://localhost --title="$title"

for plugin in $plugins
do
  if wp plugin is-installed $plugin
  then
    wp plugin activate $plugin !!!ACTIVATIONTYPE!!!
  else
    echo "\033[96mWarning:\033[0m Plugin '"$plugin"' could not be found. Have you installed it?"
  fi
done

if wp theme is-installed $theme
then
  !!!THEMEENABLE!!!
  wp theme activate $theme
else
  echo "\033[96mWarning:\033[0m Theme '"$theme"' could not be found. Have you installed it?"
fi

import() {
  for file in $content/*.xml
  do
    echo "Importing $file..."
    wp import $file --authors=skip
  done
}

if [ "$(ls -A $content)" ]
then
  if wp plugin is-installed wordpress-importer
  then
    wp plugin activate wordpress-importer
    import
  else
    echo "WordPress Importer not installed... installing now"
    wp plugin install wordpress-importer --activate
    import
    wp plugin uninstall wordpress-importer --deactivate
  fi
else
  echo "No content to be imported"
fi`)
    if multisite {
        contents = findAndReplace(contents, []byte("!!!INSTALLTYPE!!!"), []byte("multisite-install"))
        contents = findAndReplace(contents, []byte("!!!ACTIVATIONTYPE!!!"), []byte("--network"))
        contents = findAndReplace(contents, []byte("!!!THEMEENABLE!!!"), []byte("wp theme enable --network $theme"))
    } else {
        contents = findAndReplace(contents, []byte("!!!INSTALLTYPE!!!"), []byte("install"))
        contents = findAndReplace(contents, []byte("!!!ACTIVATIONTYPE!!!"), []byte(""))
        contents = findAndReplace(contents, []byte("!!!THEMEENABLE!!!"), []byte(""))
    }
    creating("setup/internal.sh", contents)
}

func makedir(directory string) {
    println("-> Creating", directory)
    os.Mkdir(directory, 0755)
}

func creatingNoexec(file string, data[]byte) {
    println( "-> Creating", file)
    ioutil.WriteFile(file, data, 0644)
}

func creating(file string, data[]byte) {
    creatingNoexec(file, data)
    os.Chmod(file, 0755)
}

func findAndReplace(haystack[]byte, old[]byte, new[]byte) []byte {
    newHaystack := bytes.Replace(haystack, old, new, -1)
    return newHaystack
}
