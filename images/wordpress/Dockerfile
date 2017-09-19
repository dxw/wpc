FROM php:apache

RUN apt-get update && apt-get install -y --no-install-recommends libpng12-dev vim-tiny less mysql-client libgmp-dev && rm -r /var/lib/apt/lists/*

RUN docker-php-ext-install mysqli && \
    docker-php-ext-install gd && \
    docker-php-ext-install gmp && \
    a2enmod rewrite

RUN curl --silent https://raw.githubusercontent.com/wp-cli/builds/gh-pages/phar/wp-cli.phar > /usr/local/bin/wp-cli.phar

ENV WORDPRESS_VERSION=latest

COPY wp /usr/local/bin/
COPY wp-start /usr/local/bin/

RUN wp core download --path=/var/www/html

COPY php.ini /usr/local/etc/php/
COPY wordpress.conf /etc/apache2/sites-enabled/
COPY mu-plugins /usr/src/mu-plugins/
COPY wp-config.php /var/www/html/wp-config.php

ENTRYPOINT ["wp-start"]
CMD ["apache2-foreground"]
