package main

const (
	EXTERNALCONTENT = `#!/bin/sh
set -e

##
## This code will be run during setup, OUTSIDE the container.
##
## Because ` + "`" + `whippet` + "`" + ` running inside the container wouldn't be able to connect
## to private repositories.
##

if test -f whippet.json; then
whippet deps install
fi`

	INTERNALCONTENT = `#!/bin/sh
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

wp core {{.InstallType}} --skip-email --admin_user=admin --admin_password=admin --admin_email=admin@localhost.invalid --url=http://localhost --title="$title"

for plugin in $plugins
do
if wp plugin is-installed $plugin
then
wp plugin activate $plugin {{.ActivationType}}
else
echo "\033[96mWarning:\033[0m Plugin '"$plugin"' could not be found. Have you installed it?"
fi
done

if wp theme is-installed $theme
then
{{.ThemeEnable}}
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
fi`
)
