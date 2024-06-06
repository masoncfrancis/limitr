# Use the official Golang image as the base image for building
FROM golang:1.22.3-alpine3.20 AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the source code
COPY . .

# Build the binary, with all the dependencies included
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o limitr ./cmd/limitr

# Use a smaller base image for running the application
FROM alpine:latest

# Set the working directory
WORKDIR /root/

# Copy the binary from the builder stage
COPY --from=builder /app/limitr .

# Specify the command to run your binary
CMD ["./limitr"]