# Use the official Golang image as a base image
FROM golang:latest

# Set the working directory inside the container
WORKDIR /usr/local/go/src/chat-grpc

RUN apt-get update && \
    apt-get install -y protobuf-compiler 

RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest

COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the server source code and necessary files into the container
COPY . .

# Go inside the chat_server directory
WORKDIR /usr/local/go/src/chat-grpc/chat_server

# Build the server application
RUN go build -o chat_server .

# Command to run the server application
CMD ["./chat_server"]
