# Use the official Golang image as the base image
FROM golang:1.22.3-alpine3.20

# Set the working directory inside the container
WORKDIR /app

# Copy the source code
COPY ..

RUN go build ./cmd/limit

# Specify the command to run your binary
CMD ["./limit"]
