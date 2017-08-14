# WPC: WordPress & (Docker) Compose

## Installation

`docker-compose` must be installed.

```
git clone https://github.com/dxw/wpc
sudo cp wpc/bin/* /usr/local/bin/
```

## Setting up a project

```
wpc_init name-of-project [--multisite]
```

You can edit `setup/internal.sh` to enable plugins and themes using wp-cli. (`setup/external.sh` also exists, if you need to run commands outside of the container).

WXR files containing WordPress content can be placed in the project's `setup/content/` folder. This content will be imported on running the project for the first time.

## Usage

Running the project for the first time:

```
docker-compose up -d
./bin/setup
```

Running the project after that:

```
docker-compose up -d
```

WordPress listens on port 80. Mailcatcher listens on port 1080. Beanstalkd Console listens on port 2080.

MySQL is available by running `./bin/wp db cli`. It will also be listening on port 3306.

To access bash within the container: `./bin/console`.

## Common issues

All the ports that docker will be listening on (80, 1080, 2080, 3066 & 11300) need to be available, so if you've got local processes running on them, that will stop the containers starting up, usually with an `EADDRINUSE` error.

Things that might be running on port 80: apache (stop with `sudo apachectl stop`), any other local server (e.g. `whippet-server`, XAMPP)

Things that might be running on port 3306: `mysql` (stop with `mysql.server stop`)
