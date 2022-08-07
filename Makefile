APP_NAME=hare-cli

build:
	@go build -o ./$(APP_NAME) cmd/cli/cli.go 

install:
	@go build -o $(GOPATH)/bin/hare cmd/cli/cli.go 

uninstall:
	@rm $(GOPATH)/bin/hare

deps:
	@go mod download

test:
	@go test ./... -v

test-coverage:
	@go test -v -covermode=count -coverprofile=coverage.out ./...