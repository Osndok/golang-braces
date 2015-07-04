
SELF=github.com/Osndok/golang-braces

build:
	go build -o $(GOPATH)/bin/go-fix-braces $(SELF)/cmd/go-fix-braces

