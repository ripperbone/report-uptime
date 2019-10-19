build:
	go build

install: build
	go install

test: build
	go test -v



