# WPC: WordPress in Containers

This repo contains the [WordPress docker image](https://hub.docker.com/repository/docker/thedxw/wpc-wordpress/) dxw uses for local development.

The image includes a mailcatcher (to avoid unintentionally sending emails to real addresses), and WP CLI.

It is *not* suitable for use in production.

Any changes to the `main` branch of this repo will automatically update the image tagged with `:latest` on Docker Hub.

If you're building a new WordPress project, you should use [dxw's WordPress Template](https://github.com/dxw/wordpress-template), which uses this image in its `docker-compose.yml`.

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

## Licence

[MIT](COPYING.md)
