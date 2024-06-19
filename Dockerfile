# Use the official Golang image as the base image
FROM golang:1.21

# Set the working directory inside the container
WORKDIR /app

# Copy the go.mod file to the working directory
COPY go.mod .

# Copy the main.go file to the working directory
COPY main.go .

# Copy the entire golangsidang directory to the working directory
COPY controllers/ ./controllers/
COPY helpers/ ./helpers/
COPY golangsidang/ ./golangsidang/
COPY middleware/ ./middleware/
COPY models/ ./models/
COPY routers/ ./routers/
COPY utils/ ./utils/

# Download dependencies defined in go.mod
RUN go get

# Perform any additional tidy up of dependencies
RUN go mod tidy

# Build the Go application and output the binary to the bin directory
RUN go build -o bin .

# Expose port 3000 to the outside world
EXPOSE 3000

# Set the entry point for the container to run the compiled binary
ENTRYPOINT [ "/app/bin" ]

# Set the default command for the container to run the API server
CMD ["./main"]