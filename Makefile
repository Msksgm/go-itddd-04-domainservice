up:
	docker compose up -d
down:
	docker compose down --remove-orphans
destroy:
	docker compose down --rmi all --volumes --remove-orphans
create-migration:
	docker compose exec app ./migrate create -ext sql -dir migrations -seq ${file}
run:
	docker compose exec app go run main.go
test:
	docker compose exec app go test ./...
