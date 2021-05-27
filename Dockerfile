FROM golang:1.15

WORKDIR /bookstore
COPY . .

RUN go mod download
RUN CGO_ENABLED=0 go build -o ./build/api ./cmd/main.go

FROM alpine:latest

WORKDIR /bookstore
COPY --from=0 /bookstore .

RUN apk add --no-cache postgresql-client
RUN chmod +x wait-for-postgres.sh

CMD ["./build/api"]