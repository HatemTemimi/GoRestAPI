serve:
	docker stop postgres-container
	sleep 1
	docker run -d -e TZ=UTC -p 5432:5432 -e POSTGRES_USER=SuperUser -e POSTGRES_PASSWORD=SuperSecure ubuntu/postgres:14-22.04_beta
	go run main.go