# Use a slim Golang base image
FROM golang:alpine

# Work directory inside the container
WORKDIR /app

# Copy the Go project files
COPY . .

# Initialize a Go module (if it doesn't exist)
RUN go mod init hello-docker-go

# Download dependencies
RUN go mod download

# Compile the application
RUN go build -o app

# Expose port 8080 of the container 
EXPOSE 8080

# Command to run when the container starts
CMD ["./app"] 
