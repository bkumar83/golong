# Stage 1: Build Go from source using Go 1.21 as a bootstrap
FROM debian:bullseye AS builder

# Install dependencies for building Go
RUN apt-get update && apt-get install -y wget tar gcc libc6-dev make

# Download and install Go 1.21 as a separate bootstrap version
RUN wget https://go.dev/dl/go1.21.1.linux-amd64.tar.gz && \
    tar -C /usr/local -xzf go1.21.1.linux-amd64.tar.gz

# Set up environment variables for Go build
ENV GOROOT_BOOTSTRAP=/usr/local/go1.21

# Download and build Go 1.24 from source
RUN wget https://go.dev/dl/go1.24.0.src.tar.gz && \
    tar -C /usr/local -xzf go1.24.0.src.tar.gz && \
    cd /usr/local/go/src && \
    ./make.bash

# Stage 2: Use the built Go binary for the final image
FROM debian:bullseye

# Copy the Go installation from the builder stage
COPY --from=builder /usr/local/go /usr/local/go

# Set environment variables for Go
ENV PATH="/usr/local/go/bin:${PATH}"
ENV CGO_ENABLED=1
ENV GOOS=linux
ENV GOARCH=amd64

# Install SQLite dependencies (for CGO)
RUN apt-get update && apt-get install -y gcc libc6-dev

# Set the working directory
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Run go mod tidy and download dependencies
RUN go mod tidy
RUN go mod download

# Copy the application files
COPY . .

# Build the application
RUN go build -o main .

# Expose the application port
EXPOSE 8080

# Run the application
CMD ["./main"]
