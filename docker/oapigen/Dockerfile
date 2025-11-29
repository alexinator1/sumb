FROM golang:1.25-alpine
RUN go install github.com/deepmap/oapi-codegen/v2/cmd/oapi-codegen@v2.2.0
WORKDIR /app
ENTRYPOINT ["/go/bin/oapi-codegen"]