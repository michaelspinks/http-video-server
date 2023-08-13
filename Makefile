
BINARY_NAME=http-video-server

build:
	@go build -o bin/${BINARY_NAME}

run: build
	@./bin/${BINARY_NAME}

clean:
	@rm -rf ./bin

test:
	@go test -v ./...