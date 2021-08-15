FROM golang:1.16 AS builder

RUN mkdir /app
ADD . /app
WORKDIR /app

RUN CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o app ./main.go

FROM alpine:latest AS production
COPY --from=builder /app .
CMD ["./app"]