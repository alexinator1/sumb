FROM golang:1.25-alpine AS builder
RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

FROM alpine:latest
COPY --from=builder /go/bin/migrate /usr/local/bin/
ENTRYPOINT ["migrate"]
