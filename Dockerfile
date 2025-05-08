# Build stage
FROM golang:1.24-alpine AS builder


# Set working directory
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./


# Download swag command line tool
RUN go install github.com/swaggo/swag/cmd/swag@latest \
    && go get -u github.com/swaggo/gin-swagger \
    && go get -u github.com/swaggo/swag \
    && go get -u github.com/swaggo/files \
    && go mod tidy

# Download dependencies
RUN go mod download
RUN go mod tidy


# Copy source code
COPY . .

# Build the application
RUN swag init -g ./cmd/http_server/main.go
RUN CGO_ENABLED=0 GOOS=linux go build -o pack_service ./cmd/http_server/main.go

# Final stage
FROM alpine:latest


WORKDIR /app

# Copy the binary from builder
COPY --from=builder /app/pack_service .

# Copy config files if needed
COPY --from=builder /app/configs ./configs

# Copy .env file
COPY --from=builder /app/.env .

# Expose port
EXPOSE 8080

# Run the application
CMD ["./pack_service"] 