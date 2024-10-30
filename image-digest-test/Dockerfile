# Use the official Go image to build the app
FROM golang:1.22 AS builder

WORKDIR /app

# Copy go.mod and go.sum files to download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY main.go .

# Build the binary with CGO disabled for compatibility with Alpine
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /digest-tester main.go

# Use a lightweight image to run the app
FROM alpine:latest
COPY --from=builder /digest-tester /digest-tester

# Run the Go app
CMD ["/digest-tester"]
