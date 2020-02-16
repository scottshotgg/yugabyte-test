FROM golang:alpine

WORKDIR /app

COPY . .
RUN go build -mod=vendor -o yugabyte-test cmd/yugabyte-test/main.go
CMD /app/yugabyte-test