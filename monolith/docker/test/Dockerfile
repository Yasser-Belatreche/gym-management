FROM golang:1.22.2

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

# Install gow for hot reloading
RUN go install github.com/mitranim/gow@latest

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["gow", "-c", "run", "main.go"]