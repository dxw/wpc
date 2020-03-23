# WPC: WordPress in Containers

This tool is used to add shell scripts and configuration files that allow running WordPress in Docker containers.

If the repository you're using already has `wpc` you'll find the following files:

- `script/server` - run WordPress in a container
  - WordPress will be running on http://localhost
  - MailCatcher will be running on http://localhost:1080
  - Beanstalk Console will be running on http://localhost:2080
- `script/setup` - install WordPress
  - the default user is `admin` with password `admin`
  - will also install [whippet](https://github.com/dxw/whippet) dependencies
  - setup configuration files are in the `setup/` directory
- `script/update` - run any updates after checking out the latest version
  - this will update whippet dependencies to the latest versions
- `script/console` - open a `/bin/sh` console inside the main container
- `script/bootstrap` - update dependencies (this is called by `script/setup` and `script/update`)
- `bin/wp` - run [wp-cli](https://wp-cli.org/)
- `docker-compose.yml` - the Docker Compose configuration for running the development environment
- `config/server.php` - configuration that would normally go in `wp-config.php` (checked into the repository - used for the development envirnoment)
- `config/server-local.php` - configuration that would normally go in `wp-config.php` (not checked into the repository)
- `setup/internal.sh` - this is run by `script/setup`, it runs the wp-cli commands to install WordPress

## Requirements

- [Docker](https://docs.docker.com/docker-for-mac/install/)
- [Docker Compose](https://docs.docker.com/compose/install/)

## Installing `wpc` onto your computer

### Method 1: With Docker

Add this alias to your shell's configuration:

```
alias wpc='docker run -ti --rm -v `pwd`:/app thedxw/wpc'
```

### Method 2: With Go

```
go get github.com/dxw/wpc
```

## Installing `wpc` into a WordPress project

### Step 1: Run `wpc`

Run `wpc name-of-project` or `wpc name-of-project --multisite` in the root of the project.

`--multisite` will add the appropriate configuration options for running the site as multisite.

### Step 2: Configuration

Edit `setup/internal.sh` to enable plugins and themes or set the name of the site, or anything else you can do with `wp-cli`. (`setup/external.sh` also exists, if you need to run commands outside of the container).

WXR files containing WordPress content can be placed in the project's `setup/content/` folder. This content will be imported on running the project for the first time.

## Running WordPress in a `wpc` project

Running it for the first time:

```
script/server
# Wait for the server to be running
script/setup
```

Running the project the next time:

```
script/server
```

Running the project after checking out the latest version:

```
script/server
# Wait for the server to be running
script/update
```

- WordPress: http://localhost
- MailCatcher: http://localhost:1080
- Beanstalk Console: http://localhost:2080
- `/bin/sh` console running on the WordPress container: `script/console`
- MySQL: `bin/wp db cli`

## Common issues

All the ports that docker will be listening on (80, 1080, 2080, 3066 & 11300) need to be available, so if you've got local processes running on them, that will stop the containers starting up, usually with an `EADDRINUSE` error.

Things that might be running on port 80: apache (stop with `sudo apachectl stop`), any other local server (e.g. XAMPP)

Things that might be running on port 3306: `mysql` (stop with `mysql.server stop`)

## WordPress versions

When the image is built, a copy of the latest version of WordPress is downloaded.

When running the image, it will check for updates before starting Apache. However it can be configured to download other versions of WordPress via the `WORDPRESS_VERSION` environment variable.

Example `docker-compose.yml` file (the last two lines should be added - the rest are there by default):

```
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
      - ./wp-content:/var/www/html/wp-content
    environment:
      WORDPRESS_VERSION: 4.7
```

## Development

Templates live in `templates/` and can be compiled into `raw_templates.go` by running `go generate`.

## Licence

[MIT](COPYING.md)
