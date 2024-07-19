# Build stage
FROM golang:1.22.5 AS builder
WORKDIR /app
COPY go.mod ./
RUN go mod tidy
COPY . /app
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o intro-golang

# Runtime stage
FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/intro-golang .
CMD ["./intro-golang"]
