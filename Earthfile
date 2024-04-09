VERSION 0.6
FROM golang:1.15-alpine3.13
WORKDIR /go-example

deps:
    COPY go.mod go.sum ./
    RUN go mod download
    # Output these back in case go mod download changes them.
    SAVE ARTIFACT go.mod AS LOCAL go.mod
    SAVE ARTIFACT go.sum AS LOCAL go.sum

build:
    FROM +deps
    COPY . .
    RUN ls -al
    RUN go build -o build/go-example main.go
    SAVE ARTIFACT build/go-example /go-example

docker:
    COPY +build/go-example .
    EXPOSE 3000
    ENTRYPOINT ["/go-example/go-example"]
    SAVE IMAGE go-example:latest