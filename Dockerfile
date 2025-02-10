# Build stage
FROM golang:1.21-alpine AS builder

WORKDIR /app

# Install git and build tools
RUN apk add --no-cache git gcc musl-dev

# Install Swagger
RUN go install github.com/swaggo/swag/cmd/swag@latest

# Set GOPATH and PATH
ENV GOPATH=/go
ENV PATH=$GOPATH/bin:$PATH

# Copy go mod and sum files
COPY go.mod ./
COPY go.sum ./

# Download dependencies and verify
RUN go mod download
RUN go mod verify

# Copy the source code
COPY . .

# Generate Swagger documentation
RUN swag init -g cmd/api/main.go -o docs --parseInternal --parseDependency

# Ensure all modules are downloaded
RUN go mod download all
RUN go mod tidy

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/api

# Final stage
FROM alpine:latest

WORKDIR /app

# Install CA certificates
RUN apk --no-cache add ca-certificates

# Copy binary and necessary files from builder
COPY --from=builder /app/main .
COPY --from=builder /app/.env .
COPY --from=builder /app/docs ./docs
COPY --from=builder /app/go.mod .
COPY --from=builder /app/go.sum .

# Make binary executable
RUN chmod +x main

# Expose port 8080
EXPOSE 8080

# Command to run the application
CMD ["./main"] 