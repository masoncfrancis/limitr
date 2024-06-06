# This Dockerfile works by first building the application in a container with the Golang image, and then copying the
# binary to a smaller image for running the application. The final image is based on the alpine image, which is a very
# small Linux distribution. The point of doing this is to make the final image as small as possible.

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