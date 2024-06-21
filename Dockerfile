FROM golang:alpine3.20

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o /usr/bin/report-uptime

EXPOSE 9095

CMD ["report-uptime"]
