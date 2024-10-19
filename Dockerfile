FROM golang:1.21 as builder

WORKDIR /app

COPY go.mod main.go ./

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o main .

FROM alpine:3.18 AS certs

RUN apk --no-cache add ca-certificates

FROM scratch

COPY --from=builder /app/main /
COPY --from=certs /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

EXPOSE 8080

CMD ["/main"]
