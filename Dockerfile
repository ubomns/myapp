# Use the official Go image as the base image
FROM golang:1.20-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files and download dependencies
COPY go.mod ./

# Copy the source code into the container
COPY . .

# Build the Go application
RUN go build -o /myapp

# Expose the application port
EXPOSE 8080

# Command to run the executable
CMD ["/myapp"]