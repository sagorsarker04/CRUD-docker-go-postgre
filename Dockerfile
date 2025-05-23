# Use an official Go runtime as a parent image
FROM golang:1.20-alpine

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the current directory contents into the container at /app
COPY . .
RUN go mod tidy

COPY .env .env
# Expose port 8080
EXPOSE 8080

# CMD with a simple placeholder Go application (this can be replaced with actual code later)
CMD ["go", "run", "main.go"]