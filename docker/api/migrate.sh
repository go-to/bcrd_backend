#!/bin/sh
set -e

# .envファイルを読み込み
if [ -f .env ]; then
#  export $(cat .env | grep -v '#' | awk '/=/ {print $1}')
#  migrate -path migrations -database "postgresql://${DB_USER}:${DB_PASS}@${DB_HOST}:${DB_PORT}/${DB_NAME}?search_path=bcrd&sslmode=disable" -verbose up
  make db-migrate-up
fi
