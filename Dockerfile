# Use the official Golang image to create a build artifact.
FROM golang:1.23.2 AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o health-check-alert .

# Start a new stage from scratch
FROM scratch

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/health-check-alert /app/health-check-alert

# Set environment variables
ENV URL_TO_PING=""
ENV DISCORD_WEBHOOK_URL=""
ENV PING_INTERVAL="5"

# Command to run the executable
CMD ["/app/health-check-alert"]