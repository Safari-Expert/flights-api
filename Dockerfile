# Stage 1: Build the Go application
FROM golang:1.22.3 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Build the Go application
RUN go build -o backend .


FROM debian:latest

# Install necessary libraries for running the Go application
RUN apt-get update && \
    apt-get install -y ca-certificates && \
    rm -rf /var/lib/apt/lists/*

# Create the directory for the application
RUN mkdir /app

# Copy the Go binary from the build stage
COPY --from=builder /app/backend /usr/local/bin/backend

# Set the working directory
WORKDIR /app

# Ensure the binary is executable
RUN chmod +x /usr/local/bin/backend

# Set the command to run the Go application
CMD ["backend"]