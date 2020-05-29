up:
	sudo docker-compose up --build

down:
	sudo docker-compose down -v

db_migrate:
	sudo docker exec CA_go go build -o cmd/main cmd/main.go
	sudo docker exec CA_go ./cmd/main db:migrate

db_reset:
	sudo docker exec CA_go go build -o cmd/main cmd/main.go
	sudo docker exec CA_go ./cmd/main db:reset