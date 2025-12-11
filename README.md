# backend-workbuddy-api

# Hexagonal Architecture with Go, Gin & GORM

## 📁 Project Structure

```
project-root/
├── cmd/
│   └── api/
│       └── main.go                 # Entry point
├── internal/
│   ├── core/                       # Domain Layer (Business Logic)
│   │   ├── domain/
│   │   │   └── user.go            # Domain models
│   │   ├── ports/
│   │   │   ├── input/
│   │   │   │   └── user_service.go    # Input ports (use cases)
│   │   │   └── output/
│   │   │       └── user_repository.go # Output ports (interfaces)
│   │   └── services/
│   │       └── user_service.go    # Business logic implementation
│   ├── adapters/                   # Adapters Layer
│   │   ├── input/
│   │   │   └── http/
│   │   │       ├── handler/
│   │   │       │   └── user_handler.go    # HTTP handlers
│   │   │       ├── dto/
│   │   │       │   └── user_dto.go        # Data transfer objects
│   │   │       └── router/
│   │   │           └── router.go          # Route definitions
│   │   └── output/
│   │       ├── persistence/
│   │       │   ├── postgres/
│   │       │   │   ├── user_repository.go # Repository implementation
│   │       │   │   └── database.go        # Database connection
│   │       │   └── entities/
│   │       │       └── user_entity.go     # Database entities
│   │       └── external/
│   │           └── email_service.go       # External services
│   └── config/
│       └── config.go              # Configuration
├── pkg/                            # Shared utilities
│   ├── logger/
│   │   └── logger.go
│   └── validator/
│       └── validator.go
├── go.mod
└── go.sum
```

## 🏗️ Architecture Layers

### 1. Domain Layer (Core)
- **Domain Models**: Pure business entities
- **Ports**: Interfaces defining contracts
- **Services**: Business logic implementation

### 2. Adapters Layer
- **Input Adapters**: HTTP handlers, gRPC, CLI
- **Output Adapters**: Database, external APIs, message queues

### 3. Infrastructure
- Configuration, logging, utilities

---