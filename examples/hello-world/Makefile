.PHONY: deps clean build

deps:
	go get -u ./...

clean: 
	rm -rf ./custom/hello-world
	
build:
	GOOS=linux GOARCH=amd64 go build -o custom/hello-world ./custom