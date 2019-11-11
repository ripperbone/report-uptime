
build:
	go build

all:
	env GOOS=linux GOARCH=amd64 go build -o linux/amd64/report-uptime
	env GOOS=linux GOARCH=arm go build -o linux/arm/report-uptime
	env GOOS=linux GOARCH=arm64 go build -o linux/arm64/report-uptime

install: build
	go install

test: build
	go test -v

