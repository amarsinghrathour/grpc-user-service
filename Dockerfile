# Stage 1: Build the Go binary
FROM golang:1.22-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the go.mod and go.sum files
COPY go.mod go.sum ./

# Download all the dependencies
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the Go binary
RUN go build -x -o server ./cmd/server/main.go
# Stage 2: Create a lightweight image to run the Go binary
FROM alpine:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/server .

# Expose the port on which the gRPC server will run
EXPOSE 50051

# Command to run the binary
CMD ["./server"]
