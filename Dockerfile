# Use the official Golang image as a base
FROM golang:1.20-alpine

# Set the working directory
WORKDIR /app

# Copy the Go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application code
COPY . ./

# Build the application binary
RUN go build -o main .

# Expose the port that the app will run on
EXPOSE 8080

# Run the application
CMD ["./main"]
