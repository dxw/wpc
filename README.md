# WPC: WordPress & (Docker) Compose

## Installation

`docker-compose` must be installed.

```
git clone https://github.com/dxw/wpc
sudo cp wpc/bin/* /usr/local/bin/
```

## Setting up a project

```
wpc_init name-of-project
```

This will create a `docker-compose.yml` file in the current directory.

## Usage

Running the project:

```
docker-compose up
```

WordPress will now be listening on port 80. Mailcatcher will be available on port 1080.

To access bash within the WordPress container:

```
wpc_console
```

To access wp-cli:

```
wpc_cli --info
wpc_cli db cli
wpc_cli db import - < database.sql
wpc_cli db export - > database.sql
```
