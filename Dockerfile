FROM golang:1.24.3 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go test ./...
# RUN go build -o rem ./cmd/rem/main.go
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o rem ./cmd/rem/main.go

RUN chmod +x rem

#-----------------------------------
FROM debian:bookworm AS runner
ENV TZ=Asia/Tokyo

RUN apt update \
    && apt install -y --no-install-recommends ca-certificates \
    && apt clean \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /app

COPY --from=builder /app/rem /app/rem

CMD ["./rem"]
