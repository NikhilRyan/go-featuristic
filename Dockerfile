FROM golang:latest

# Install necessary packages
RUN apt-get update && apt-get install -y \
    git \
    && rm -rf /var/lib/apt/lists/*

# Create a directory for the app
RUN mkdir -p /go/src/go-featuristic
WORKDIR /go/src/go-featuristic

# Copy the project files
COPY . .

# Download Go modules
RUN go mod tidy

# Expose application port
EXPOSE 8080
