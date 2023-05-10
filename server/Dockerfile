# Dockerfile
FROM golang:latest

# Set the working directory for the app
WORKDIR /go/src/app

# Copy the app code into the container
COPY . .

# Install dependencies and build the app
RUN go get -d -v ./...
RUN go build -o main .

# Expose port 8080 for the app
EXPOSE 8080

# Set the entrypoint for the container to run the app
CMD ["./main"]
