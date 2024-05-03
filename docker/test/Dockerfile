FROM golang:latest

# Set the Current Working Directory inside the container
WORKDIR /home/app

# Copy needed files
COPY go.mod ./
COPY go.sum ./
COPY main.go ./
COPY src ./src

# Download all dependencies.
RUN go mod download
RUN go mod verify

# Build the Go app
RUN go build -o main .

# This is the port that the application will use
ENV PORT=8080

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./main"]