# Use the official Go image as the base image
FROM golang:1.22-alpine

# Set the working directory inside the container
WORKDIR /goapp
#COPY ./ /goapp

# Copy the Go module files
COPY go.mod go.sum /goapp/

# Download the dependencies
RUN go mod download

# Copy the rest of the application code
COPY . /goapp
#COPY .env /goapp

# Build the Go application
RUN go build -o main .

# Expose the application on port 8080
EXPOSE 8080

# Set the entry point command to run the built binary
#CMD ["./main"]
CMD ["/usr/local/go/bin/go", "run", "main.go"]
