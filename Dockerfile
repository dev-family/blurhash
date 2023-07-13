FROM golang:1.18-alpine AS builder

WORKDIR /app

COPY main.go go.mod ./
RUN go mod download

COPY go.mod go.sum main.go ./

RUN apk add --no-cache ca-certificates
RUN apk add --no-cache git
RUN CGO_ENABLED=0 GOOS=linux go build -o blurhash

FROM scratch

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /app/blurhash /

CMD ["/blurhash"]
