install:
	go get -u github.com/gin-gonic/git
	go get -u github.com/stretchr/testify

run:
	@go run main.go

build:
	@go build

test:
	@cd tests && go test

