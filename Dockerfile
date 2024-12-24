# Stage 1: Build stage
FROM golang:1.22.2-alpine AS builder

# Set environment variables for Go
ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Create a working directory
WORKDIR /server

# Copy go mod and sum files
COPY go.mod ./

# Copy the source code
COPY . ./

# Build the application
RUN go build -o server .

# Stage 2: Production stage
FROM alpine:latest

# Create a working directory
WORKDIR /root/

# Copy the compiled binary from the builder stage
COPY --from=builder /server .

# Command to run the application
CMD ["./server"]
