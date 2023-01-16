# Source code build
FROM golang:1.14 AS builder
WORKDIR /app
COPY ./main.go .
RUN go build -o sauth .

# Production enviroment setup
FROM debian:9

WORKDIR /var/lib/sauth
EXPOSE 1930

COPY --from=builder /app/sauth /usr/local/bin/sauth

CMD ["/usr/local/bin/sauth"]