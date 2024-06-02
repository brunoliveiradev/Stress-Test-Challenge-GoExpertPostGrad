FROM golang:latest as builder

LABEL authors="brunooliveira"

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -o /app/main ./cmd/stress-test

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/main .

ENTRYPOINT ["./main"]
