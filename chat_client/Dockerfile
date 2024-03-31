# Use the official Golang image as a base image
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the client source code and necessary files into the container
COPY . .

# Build the client application
RUN go build -o client .

# Command to run the client application
CMD ["./client"]
