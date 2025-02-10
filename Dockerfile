# Build stage
FROM golang:1.21-alpine AS builder

WORKDIR /app

# Install git and build tools
RUN apk add --no-cache git gcc musl-dev

# Copy go mod and sum files
COPY go.mod ./

# Initialize go.mod and download dependencies
RUN go mod download
RUN go mod tidy

# Create go.sum if it doesn't exist
RUN touch go.sum

# Copy the source code
COPY . .

# Download all dependencies again to ensure go.sum is up to date
RUN go mod download
RUN go mod tidy
RUN go mod verify

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o /go/bin/app ./cmd/api

# Final stage
FROM alpine:latest

# Install CA certificates and create app directory
RUN apk --no-cache add ca-certificates && \
    mkdir -p /app

WORKDIR /app

# Copy binary and config from builder
COPY --from=builder /go/bin/app /app/app
COPY --from=builder /app/.env /app/.env

# Make the binary executable
RUN chmod +x /app/app

# Expose port 8080
EXPOSE 8080

# Command to run the application
CMD ["/app/app"] 