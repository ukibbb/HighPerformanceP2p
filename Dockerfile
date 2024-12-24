FROM golang:1.22.2-alpine AS builder

ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /root

COPY go.mod ./

COPY . ./

RUN go build -o server .

# Stage 2: Production stage
FROM alpine:latest

# Create a working directory
WORKDIR /root/

# Copy the compiled binary from the builder stage
COPY --from=builder /root/server ./server

# Command to run the application
CMD ["./server"]
