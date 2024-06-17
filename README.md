# Go Fiber Todo API

This project is a RESTful API for managing todo items, built with Go and the Fiber web framework. It follows clean architecture principles and incorporates dependency injection for improved modularity and testability.

## Code Style

This project follows the guidelines outlined in [Effective Go](https://go.dev/doc/effective_go). Please ensure your contributions adhere to these standards.

## Project Structure

```
.
├── cmd
│   └── api
│       └── main.go
├── coverage.out
├── go.mod
├── go.sum
├── internal
│   ├── app
│   │   └── app.go
│   ├── configs
│   │   └── config.go
│   ├── handlers
│   │   ├── todo_handler.go
│   │   └── todo_handler_test.go
│   ├── middlewares
│   │   └── todo_middleware.go
│   ├── models
│   │   └── todo_model.go
│   ├── repositories
│   │   ├── todo_repository.go
│   │   └── todo_repository_test.go
│   ├── routes
│   │   └── todo_route.go
│   └── services
│       ├── todo_service.go
│       └── todo_service_test.go
├── Makefile
└── pkg
    ├── databases
    │   └── mongo.go
    └── di
        └── todo_di.go
```

## Features

- CRUD operations for todo items
- Clean architecture
- Dependency Injection
- Unit tests
- MongoDB integration
- Middleware support
- Logging

## Getting Started

1. Clone the repository:
   ```
   git clone https://github.com/ferhateroglu/go-fiber.git
   ```

2. Navigate to the project directory:
   ```
   cd go-fiber
   ```

3. Install dependencies:
   ```
   go mod tidy
   ```

4. Set up your MongoDB connection in `.env`.

5. Run the application:
   ```
   make run
   ```

## Testing

Run the tests using the following command:

```
make test
```

To generate a coverage report, run:

```
make coverage
```

## Project Components

- `cmd/api/main.go`: Entry point of the application
- `internal/app`: Application setup and configuration
- `internal/configs`: Configuration management
- `internal/handlers`: HTTP request handlers
- `internal/middlewares`: Custom middleware functions
- `internal/models`: Data models
- `internal/repositories`: Data access layer
- `internal/routes`: API route definitions
- `internal/services`: Business logic layer
- `pkg/databases`: Database connection setup
- `pkg/di`: Dependency injection setup

## Coverage Report

- `internal/handlers:` 77.8% of statements
- `internal/repositories:` 84.3% of statements
- `internal/services:` 96.0% of statements
