FROM golang:1.25.9-alpine3.23 AS builder

WORKDIR /app

COPY go.mod go.sum* /app/

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o pinggu .

FROM alpine:3.24

WORKDIR /app

COPY --from=builder /app/pinggu .

USER 1001

ENTRYPOINT [ "./pinggu" ]
