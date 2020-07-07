#!/bin/sh
set -e

FILES="${0} images/wordpress/wp-start images/wordpress/wp templates/bin/* templates/script/* templates/setup/*.sh.tmpl"

for I in ${FILES}; do
  echo "Checking ${I}..."
  perl -pe 's/\{\{.*?\}\}/TEMPLATE_VALUE/g' < "${I}" | shellcheck -
done

echo OK
