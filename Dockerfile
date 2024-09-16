# Stage 1: Build the Go application
FROM golang:1.23.1-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum for dependency resolution
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the application
RUN go build -o main .

# Stage 2: Create a minimal image for running the application
FROM alpine:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the built binary from the builder stage
COPY --from=builder /app/main .


EXPOSE 8080
# Command to run the application
CMD ["./main"]
