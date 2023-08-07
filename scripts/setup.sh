#!/bin/sh
DB_USER=postgres
DB_PASS=root
DB_NAME=product
DB_PORT=5432

(echo "=========Launch Start===========" & sleep 2)
#set env
rm -rf $(pwd)/app.env || true
echo DB_USER=$DB_USER >> $(pwd)/app.env
echo DB_PASS=$DB_PASS >> $(pwd)/app.env
echo DB_PORT=$DB_PORT >> $(pwd)/app.env
echo DB_NAME=$DB_NAME >> $(pwd)/app.env
#kill activity on db port
npm list -g npx || npm install -g npx
npx kill-port $DB_PORT || true
docker stop postgres || true
docker run -d  --rm --name postgres  -e TZ=UTC -p $DB_PORT:5432 -e POSTGRES_USER=$DB_USER -e POSTGRES_PASSWORD=$DB_PASS ubuntu/postgres:14-22.04_beta
(echo "=========Loading database container===========" & sleep 2)
go run main.go



