.PHONY: up
up:
	#go run cmd/smart-sr/main.go
	docker-compose up -d
	docker-compose ps

.PHONY: down
down:
	docker-compose down
	docker-compose ps

.PHONY: restart
restart:
	docker-compose down
	docker-compose up -d
	docker-compose ps

.PHONY: log
log:
	docker-compose logs -f

.PHONY: migration
migration:
	#go run cmd/migration/main.go
	docker-compose exec app go run cmd/migration/main.go

.PHONY: test
test:
	#go test -v ./...
	docker-compose exec app go test -v ./...

.PHONY: db
db:
	docker-compose exec db mysql -uroot -proot smart-sr