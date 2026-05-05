.PHONY: run test build

# Путь к конфигу
CONFIG_PATH := config/local.yaml

run:
	go run cmd/url-shortener/main.go --config=$(CONFIG_PATH)

test:
	go test -v ./... --config=$(CONFIG_PATH)

build:
	go build -o bin/url-shortener cmd/url-shortener/main.go

# Запуск с флагами
dev:
	go run cmd/url-shortener/main.go --config=config/dev.yaml

prod:
	go run cmd/url-shortener/main.go --config=config/prod.yaml