# Use the official Golang image as the base image
FROM golang:1.23-alpine

# Set environment variables for Go cross-compiling
ENV CGO_ENABLED=0
ENV GO111MODULE=on

# Install necessary dependencies (you can add more if needed)
RUN apk add --no-cache bash

# Set the working directory inside the container
WORKDIR /app

# Copy the Go source code into the container
COPY . /app

# Copy the build script into the container
COPY build.sh /app/build.sh

# Make the build script executable
RUN chmod +x /app/build.sh

# Set the default command to run the build script
CMD ["/app/build.sh"]
