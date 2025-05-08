# Pack Service

A Go-based microservice for managing and processing pack-related operations. This service is built using hexagonal architecture principles and follows best practices for Go development.

## 🚀 Features

- RESTful API endpoints for pack management
- Redis integration for caching and data storage
- Swagger documentation for API endpoints
- Structured logging with Zap
- Environment-based configuration
- hexagonal Architecture implementation

## 🛠️ Tech Stack

- **Language**: Go 1.24.1
- **Web Framework**: Gin
- **Database/Cache**: Redis
- **Documentation**: Swagger
- **Logging**: Zap
- **Configuration**: godotenv

## 📁 Project Structure

```
.
├── cmd/            # Application entry points
├── configs/        # Configuration files
├── docs/          # Documentation
├── internal/      # Private application code
│   ├── core/     # Core business logic
│   │   ├── consts/    # Constants
│   │   ├── domain/    # Domain models
│   │   ├── ports/     # Interface definitions
│   │   └── services/  # Business logic implementation
│   └── driver/   # External interfaces implementation
├── pkg/           # Public libraries
└── go.mod         # Go module definition
```

## 🚀 Getting Started

### Prerequisites

- Go 1.24.1 or higher
- Redis server
- Make (optional, for using Makefile commands)

### Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/mahdiZarepoor/pack_service_assignment.git
   cd pack_service_assignment
   ```

2. Install dependencies:
   ```bash
   go mod download
   ```

3. Set up environment variables:
   ```bash
   cp .env.example .env
   # Edit .env with your configuration
   ```

4. Run the application:
   ```bash
   go run cmd/api/main.go
   ```

## 📚 API Documentation

The API documentation is available through Swagger UI. Once the application is running, you can access it at:

```
http://localhost:8080/swagger/index.html
```

## 🧪 Testing

To run the tests:

```bash
go test ./...
```

## 🔧 Configuration

The service can be configured through environment variables. Key configuration options include:

- `PORT`: Server port (default: 8080)
- `REDIS_HOST`: Redis host address
- `REDIS_PORT`: Redis port
- `REDIS_PASSWORD`: Redis password (if required)


## 👥 Authors

- Mahdi Zarepoor - Initial work