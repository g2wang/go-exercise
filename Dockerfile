FROM golang:1.18-alpine

# Install git
RUN set -ex; \
    apk update; \
    apk add --no-cache git

# Set working directory
WORKDIR /go/src/github.com/g2wang/go-exercise
# Run tests
CMD CGO_ENABLED=0 go test -v ./orgaccnt
