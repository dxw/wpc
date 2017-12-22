package main

const (
	DOCKERCOMPOSECONTENT = `version: "3"

volumes:
mysql_data_{{.Name}}:

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
  - mysql_data_{{.Name}}:/var/lib/mysql
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
  - ./wp-content:/var/www/html/wp-content`
)
