# Use an official Golang image as the base image
FROM golang:1.20 as builder

# Set the current working directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the working directory inside the container
COPY . .

# Build the Go app and place the binary in /app/bin
RUN mkdir -p /app/bin && go build -o /app/bin/main .

# Debugging step: List the files in the /app/bin directory to ensure main is created
RUN ls -l /app/bin

# Start a new stage from scratch
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /app

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/bin/main /app/bin/main
COPY --from=builder /app/config /app/config

# Debugging step: Ensure the main binary exists and has correct permissions
RUN ls -l /app/bin && chmod +x /app/bin/main

# Command to run the executable
CMD ["/app/bin/main"]
