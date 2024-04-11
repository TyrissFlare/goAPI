# Use the official Golang image as a parent image
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files and download dependencies
COPY go.mod ./
RUN go mod download

# Copy the rest of the application source code into the container
COPY . .

# Build the Go application inside the container
RUN go build -o main .

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./main"]
