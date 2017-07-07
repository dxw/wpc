FROM php:apache

## Basic setup

# mysqli
RUN docker-php-ext-install mysqli
# gd
RUN apt-get update && apt-get install -y --no-install-recommends libpng12-dev && rm -r /var/lib/apt/lists/*
RUN docker-php-ext-install gd
# mod_rewrite
RUN a2enmod rewrite

## Main setup

RUN curl --silent https://wordpress.org/latest.tar.gz > /usr/src/latest.tar.gz \
    && tar -C /usr/src -xzf /usr/src/latest.tar.gz \
    && rm -rf /var/www/html /usr/src/wordpress/wp-content \
    && mv /usr/src/wordpress /var/www/html
COPY php.ini /usr/local/etc/php/
COPY wordpress.conf /etc/apache2/sites-enabled/
COPY mu-plugins /usr/src/mu-plugins/
COPY wp-config.php /var/www/html/wp-config.php

## Extra tools

# vi
RUN apt-get update && apt-get install -y --no-install-recommends vim-tiny && rm -r /var/lib/apt/lists/*
# wp-cli
RUN apt-get update && apt-get install -y --no-install-recommends less mysql-client && rm -r /var/lib/apt/lists/*
RUN curl https://raw.githubusercontent.com/wp-cli/builds/gh-pages/phar/wp-cli.phar > /usr/local/bin/wp-cli.phar
RUN /bin/echo -e '#!/bin/sh\nexec php /usr/local/bin/wp-cli.phar --allow-root ${@}' > /usr/local/bin/wp
RUN chmod 755 /usr/local/bin/wp
