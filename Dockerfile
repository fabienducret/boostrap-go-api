FROM golang:1.23.3-alpine as builder

RUN mkdir /app
COPY . /app
WORKDIR /app

RUN CGO_ENABLED=0 go build -o api ./cmd/api

RUN chmod +x /app/api

# build a tiny docker image
FROM alpine:latest

RUN mkdir /app

COPY --from=builder /app/api /app

CMD ["/app/api"]