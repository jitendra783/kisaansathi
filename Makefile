MAIN_FILE_PATH=app/main.go
BIN_PATH=app/kisaanSathi

default: build run

# Need to have GCC compiler in PATH or must specify
build-prod:
	CGO_ENABLED=1 GOARCH=amd64 GOOS=linux go build -tags with_appd -ldflags "-s -w -r -Wl,-rpath,libs" -o $(BIN_PATH) $(MAIN_FILE_PATH)

# For testing purpose, build the binary for prod server on macos
build-prod-macos:
	CC=x86_64-unknown-linux-gnu-gcc CGO_ENABLED=1 GOARCH=amd64 GOOS=linux go build -tags with_appd -ldflags "-s -w -r -Wl,-rpath,libs" -o $(BIN_PATH) $(MAIN_FILE_PATH)

build:
	go build -o $(BIN_PATH) $(MAIN_FILE_PATH)

run:
	./$(BIN_PATH)
