FROM php:8.2-apache

RUN apt-get update && apt-get install -y --no-install-recommends clamav libmagickwand-dev libjpeg-dev libpng-dev vim-tiny less mariadb-client libgmp-dev libsodium-dev zlib1g-dev libzip-dev && rm -r /var/lib/apt/lists/* && freshclam

RUN pecl install imagick

RUN docker-php-ext-install gd && \
    docker-php-ext-install gmp && \
    docker-php-ext-install sodium && \
    docker-php-ext-install zip && \
    docker-php-ext-install pcntl &&\
    docker-php-ext-install sockets &&\
    docker-php-ext-install sysvsem &&\
    docker-php-ext-install bcmath &&\
    docker-php-ext-install exif &&\
    docker-php-ext-install soap &&\
    docker-php-ext-enable imagick && \
    docker-php-ext-install pdo pdo_mysql &&\
    docker-php-ext-install mysqli && \
    docker-php-ext-enable mysqli && \
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

COPY openssl.cnf /etc/ssl/openssl.cnf
RUN openssl req -x509 -days 3650 -out /etc/ssl/certs/apache-localhost.crt \
    -keyout /etc/ssl/private/apache-localhost.key \
    -newkey rsa:2048 -nodes -sha256 \
    -subj '/CN=localhost' \
    -extensions EXT \
    -config /etc/ssl/openssl.cnf 


COPY apache-https.conf /etc/apache2/sites-available/
RUN a2enmod ssl
RUN a2ensite apache-https

ENTRYPOINT ["wp-start"]
CMD ["apache2-foreground"]
