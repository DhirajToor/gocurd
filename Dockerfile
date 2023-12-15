# Stage 1: Build
FROM golang:1.21-alpine AS builder

WORKDIR /app

# Install gcc and other necessary build tools
RUN apk add --no-cache build-base

# Copy go.mod and go.sum files to download dependencies
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code to the container
COPY . .

# Build the Go application
RUN CGO_ENABLED=1 go build -o gocurd

# Stage 2: Create a lightweight image
FROM alpine:latest


WORKDIR /app

# Copy only the necessary files from the builder stage
COPY --from=builder /app/gocurd .
COPY --from=builder /app/data.db .

# Expose the port the app runs on
EXPOSE 8080

# Command to run the executable
CMD ["./gocurd"]