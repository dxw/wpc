#!/bin/sh
set -e

check() {
  echo "Checking $1..."
  perl -pe 's/\{\{.*?\}\}/TEMPLATE_VALUE/g' < "$1" | shellcheck -
}

check "$0"

find images -type f \( -name wp-start -o -name wp \) -exec sh -c '
  for file do
    echo "Checking $file..."
    perl -pe "s/\{\{.*?\}\}/TEMPLATE_VALUE/g" < "$file" | shellcheck -
  done
' sh {} +

echo OK
