build: 
	go build

test:
	go test ./... -coverprofile cover.out

help: test build
	./gh-cli-go -h
	