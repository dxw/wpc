FROM php:apache

# APT
RUN apt-get update && apt-get install -y --no-install-recommends libpng12-dev vim-tiny less mysql-client && rm -r /var/lib/apt/lists/*

# PHP/Apache setup
RUN docker-php-ext-install mysqli && \
    docker-php-ext-install gd && \
    a2enmod rewrite

# wp-cli
RUN curl --silent https://raw.githubusercontent.com/wp-cli/builds/gh-pages/phar/wp-cli.phar > /usr/local/bin/wp-cli.phar

# WordPress
RUN wp core download --path=/var/www/html

COPY php.ini /usr/local/etc/php/
COPY wordpress.conf /etc/apache2/sites-enabled/
COPY mu-plugins /usr/src/mu-plugins/
COPY wp-config.php /var/www/html/wp-config.php
COPY wp /usr/local/bin/
