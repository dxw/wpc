#!/bin/sh
set -e

FILES="${0} images/wordpress/php7.4/wp-start images/wordpress/php7.4/wp"

for I in ${FILES}; do
  echo "Checking ${I}..."
  perl -pe 's/\{\{.*?\}\}/TEMPLATE_VALUE/g' < "${I}" | shellcheck -
done

echo OK
