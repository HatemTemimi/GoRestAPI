#!/bin/sh
DB_USER=postgres
DB_PASS=root
DB_NAME=productsapi
DB_PORT=5432

(echo "Launch start.." & sleep 2)
(echo "Installing dependencies..")
go install
(echo "Setting Environement.." & sleep 2)
rm -rf $(pwd)/.env || true
echo DB_USER=$DB_USER >> $(pwd)/.env
echo DB_PASS=$DB_PASS >> $(pwd)/.env
echo DB_PORT=$DB_PORT >> $(pwd)/.env
echo DB_NAME=$DB_NAME >> $(pwd)/.env
(echo "Cleaning ports.." & sleep 2)
npm list -g npx || npm install -g npx
npx kill-port $DB_PORT || true
docker stop postgres || true
docker run -d  --rm --name postgres  -e TZ=UTC -p $DB_PORT:5432 -e POSTGRES_USER=$DB_USER -e POSTGRES_PASSWORD=$DB_PASS ubuntu/postgres:14-22.04_beta
(echo "Loading pg container.." & sleep 2)
go run main.go



