# Use the official Go image as the base image
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the Go modules and other necessary files
COPY go.mod go.sum ./


# Download and install dependencies
RUN go mod download

# Copy the rest of the application files
COPY . .

# RUN curl -sSf https://atlasgo.sh | sh

# RUN atlas migrate diff --env gorm

# RUN atlas schema apply --env gorm -u "postgres://postgres:123@localhost:5432/dev-db"

WORKDIR /app/apps/Tenant

# Build the Go application
RUN go build -o tenant

# Command to run the application
CMD ["./tenant"]
