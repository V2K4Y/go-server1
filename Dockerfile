# Use the official Go image as the base image
FROM golang:1.18-alpine

# Set the working directory to /app
WORKDIR /app

# Copy the Go server source code to the container
COPY . .

# Build the Go server binary
RUN go build -o server .

# Expose port 8000 for the Go server
EXPOSE 8000

# Start the Go server when the container starts
CMD ["./server"]