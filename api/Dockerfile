# Build Stage
FROM golang:1.25-alpine AS builder

WORKDIR /opt/api

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags="-s -w" -o api .

FROM alpine:latest

WORKDIR /opt/api
COPY --from=builder /opt/api .
EXPOSE 1323

CMD ["./api"]