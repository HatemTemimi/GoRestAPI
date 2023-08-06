ifneq (,$(wildcard ./app.env))
    include app.env
    export
endif

serve:
	npx kill-port 5432
	docker stop postgres || true
	docker rm postgres || true
	docker run -d  --rm --name postgres  -e TZ=UTC -p $(DB_PORT):5432 -e POSTGRES_USER=$(DB_USER) -e POSTGRES_PASSWORD=$(DB_PASS) ubuntu/postgres:14-22.04_beta
	sleep 3
	go run main.go
	