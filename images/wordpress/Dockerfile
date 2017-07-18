FROM php:apache

# APT
RUN apt-get update && apt-get install -y --no-install-recommends libpng12-dev vim-tiny less mysql-client && rm -r /var/lib/apt/lists/*

# PHP/Apache setup
RUN docker-php-ext-install mysqli && \
    docker-php-ext-install gd && \
    a2enmod rewrite

# wp-cli
RUN curl https://raw.githubusercontent.com/wp-cli/builds/gh-pages/phar/wp-cli.phar > /usr/local/bin/wp-cli.phar &&\
    /bin/echo -e '#!/bin/sh\nexec php /usr/local/bin/wp-cli.phar --allow-root "${@}"' > /usr/local/bin/wp &&\
    chmod 755 /usr/local/bin/wp

# WordPress
RUN curl --silent https://wordpress.org/latest.tar.gz > /usr/src/latest.tar.gz && \
    tar -C /usr/src -xzf /usr/src/latest.tar.gz && \
    rm -rf /var/www/html /usr/src/wordpress/wp-content && \
    mv /usr/src/wordpress /var/www/html

COPY php.ini /usr/local/etc/php/
COPY wordpress.conf /etc/apache2/sites-enabled/
COPY mu-plugins /usr/src/mu-plugins/
COPY wp-config.php /var/www/html/wp-config.php
