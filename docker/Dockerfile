# Stage 1: Building the Go binary
FROM golang:1.22.3-alpine AS builder

WORKDIR /app

# Copy go.mod and go.sum files
COPY ../go.mod go.sum ./

# Downloading dependencies
RUN go mod download

# Copy the source code
COPY .. .

# Building the Go binary
RUN go build -o main main.go

# Stage 2: Creating a smaller image with only the necessary artifacts
FROM scratch

WORKDIR /app

# Copy the Go binary from the builder stage
COPY --from=builder /app/main .
# Copy the template files
COPY --from=builder /app/index.gohtml .

# Expose the application port
EXPOSE 8080

# Command to run the application
CMD ["./main"]
