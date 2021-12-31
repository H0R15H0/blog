#!/bin/sh

set -e

echo "Waiting for mysql"

until mysql -h"$DB_HOST" -P3306 -u"$DB_USER" -p"$DB_PASS" "$DB_NAME" --execute "SHOW DATABASES;" &> /dev/null
do
  printf "."
  sleep 1
done

  
echo "MySQL is up - executing command"
exec "$@"
