# Start from the official Go image as the build stage
FROM golang:1.20-alpine AS builder

# Set environment variables
ENV GO111MODULE=on
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

# Create and set the working directory
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the Go application
RUN go build -o makerchecker .

# Start a new stage from scratch
FROM alpine:latest

# Set the working directory
WORKDIR /root/

# Copy the binary from the builder stage
COPY --from=builder /app/makerchecker .

# Expose port 8080
EXPOSE 8080

# Command to run the executable
CMD ["./makerchecker"]
