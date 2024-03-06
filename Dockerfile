# Use the official Go image as the base image
FROM golang:1.22-alpine

# Set the working directory inside the container
WORKDIR /goapp

# Copy the Go module files
COPY go.mod go.sum /goapp/

# Download the dependencies
RUN go mod download

# Copy the rest of the application code
COPY . ./

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -o /docker-gs-ping


# Expose the application on port 8080
EXPOSE 8080

# Set the entry point command to run the built binary
CMD ["/docker-gs-ping"]
