# Use the official Go image as the base image
FROM golang:latest

RUN apt-get update && \
    apt-get install -y docker.io

RUN systemctl start docker

# Set the working directory inside the container
WORKDIR /app

# Copy the Go modules and other necessary files
COPY go.mod go.sum ./


# Download and install dependencies
RUN go mod download

# Copy the rest of the application files
COPY . .

WORKDIR /app/apps/Auth

# Build the Go application
RUN go build -o auth

# Command to run the application
CMD ["./auth"]
