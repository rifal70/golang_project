# Use the official Golang image
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the local code to the container
COPY . .

# Initialize Go modules and download dependencies
RUN go mod download

# Install the required dependencies
RUN go get -u github.com/gorilla/mux

# Build the Go application
RUN go build -o main .

# Expose the port the app runs on
EXPOSE 8069

# Command to run the executable
CMD ["./main"]
