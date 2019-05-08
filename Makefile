BINARY_NAME=test-hello

all: build
build:
	go build -o $(BINARY_NAME) -v
clean:
	go clean
	rm -f $(BINARY_NAME)
get:
	go get -d -t ./...
run:
	clear
	go build -o $(BINARY_NAME) -v
	./$(BINARY_NAME)

.PHONY: all build clean get run