.PHONY: build run test clean

# Build the application
build:
	go build -o pack_service ./cmd/app/main.go

# Run the application
run:
	go run ./cmd/app/main.go

# Run tests
test:
	go test -v ./...

# Clean build artifacts
clean:
	go clean
	rm -f pack_service

# Show help
help:
	@echo "Available commands:"
	@echo "  make build  - Build the application"
	@echo "  make run    - Run the application"
	@echo "  make test   - Run tests"
	@echo "  make clean  - Clean build files"
	@echo "  make help   - Show this help message" 