# Use the official Golang image as the builder
FROM golang:1.23-alpine AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the Go source code
COPY . .

# Build the Go application
RUN go build -o multistage

# Use a minimal base image to reduce the final image size
FROM alpine:latest

# Set the working directory
WORKDIR /root/

# Copy the built Go binary from the builder stage
COPY --from=builder /app/multistage .

# Expose port if needed (optional)
#EXPOSE 8080

# Command to run the executable
CMD ["./multistage"]
