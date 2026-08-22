APP_NAME := kisansathi-backend
MAIN_FILE := ./app/main.go

.PHONY: run build test clean docker-build docker-run docker-stop tidy

run:
	go run $(MAIN_FILE)

build:
	go build -o bin/$(APP_NAME) $(MAIN_FILE)

test:
	go test ./...

tidy:
	go mod tidy

clean:
	rm -rf bin

docker-build:
	docker build -t $(APP_NAME):latest .

docker-run:
	docker run --rm -p 8008:8008 --name $(APP_NAME) $(APP_NAME):latest

docker-stop:
	docker stop $(APP_NAME)

docker-clean:
	docker rm -f $(APP_NAME) 2>/dev/null || true
