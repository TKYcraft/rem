FROM golang:1.24.3 as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go test ./...
# RUN go build -o rem ./cmd/rem/main.go
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o rem ./cmd/rem/main.go

RUN chmod +x rem

#-----------------------------------
FROM alpine:3.21.3
ENV TZ=Asia/Tokyo

WORKDIR /app

COPY --from=builder /app/rem /app/rem

CMD ["./rem"]
