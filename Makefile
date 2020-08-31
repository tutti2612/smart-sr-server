.PHONY: run
run:
	go run cmd/smart-sr/main.go

.PHONY: migration
migration:
	go run cmd/migration/main.go

.PHONY: test
test:
	go test -v ./...

.PHONY: db
db:
	docker-compose exec db mysql -usmart-sr -psmart-sr smart-sr