FROM golang:latest AS builder
WORKDIR /app
COPY . .
RUN GOOS=linux CGO_ENABLED=0 go build -ldflags="-w -s" -o server cmd/gobooks/main.go

FROM scratch
COPY --from=builder /app/server .
CMD ["./server"]