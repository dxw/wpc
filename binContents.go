package main

func wpContents() []byte {

	return []byte(
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
}

func consoleContents() []byte {
	return []byte(
		`#!/bin/sh
set -e

exec docker-compose exec wordpress bash`)
}

func setupContents() []byte {
	return []byte(
		`#!/bin/sh
set -e
#
# Runs all site setup scripts

` + "`" + `dirname $0` + "`" + `/../setup/external.sh
docker-compose exec wordpress /usr/src/app/setup/internal.sh`)
}
