# Auth Service

A lightweight authentication microservice written in Go. The service will handle user authentication and authorization within a microservice architecture, communicating with other services via gRPC.

> **Project Status:** Early development (pre-alpha) — no stable API yet.

## Features

Current features:

* Environment-based configuration
* Structured logging
* Unit tests for configuration loading
* Reusable testing utilities
* Grpc api

Planned features:

* User registration
* User authentication
* JWT access and refresh tokens
* Password hashing
* Email verification
* PostgreSQL persistence
* Docker support
* Jenkins CI pipeline

---

## Technology Stack

* Go
* gRPC
* PostgreSQL (planned)
* GitHub Actions
* Jenkins (planned)

---

## Project Structure

```text
.
├── cmd/                # Application entry points
├── internal/
│   ├── config/         # Configuration loading
│   ├── logger/         # Logging package
│   ├── server/         # gRPC server (planned)
│   └── testutil/       # Testing helpers
├── .github/
│   └── workflows/      # GitHub Actions
├── go.mod
└── README.md
```

---

## Requirements

* Go 1.25 or newer

---

## Getting Started

Clone the repository:

```bash
git clone https://github.com/abssh/auth-service.git
cd auth-service
```

Download dependencies:

```bash
go mod download
```

Create the required environment variables.

Example:

```env
HTTP_HOST=localhost
HTTP_PORT=7000
GRPC_HOST=localhost
HTTP_PORT=7001
LOG_LEVEL=INFO
```

Run the service:

```bash
go run ./cmd/auth-service/main.go
```

---

## Running Tests

Run all unit tests:

```bash
go test ./...
```

Run tests with verbose output:

```bash
go test -v ./...
```

---

## Configuration

| Environment Variable | Description   | Default     |
| --------------------- | -------------- | ------------ |
| `HTTP_HOST`            | HTTP Server host    | Networking Deault  |
| `HTTP_PORT`            | HTTP Server port    | `8080`             |
| `GRPC_HOST`            | GRPC Server host    | Networking Deault  |
| `GRPC_PORT`            | GRPC Server port    | `9090`             |
| `LOG_LEVEL`            | Logging level       | `INFO`             |

---

## Development

CI runs via GitHub Actions on every push and pull request.

The project follows a feature branch workflow using `develop` as the integration branch.

Typical workflow:

1. Create a branch from `develop`.
2. Implement the feature.
3. Add or update tests.
4. Open a pull request targeting `develop`.
5. Ensure all CI checks pass.
6. Merge after review.

Example branch names:

```text
feature/grpc
feature/jwt
test/config
ci/github-actions
```

---

## Roadmap

* [x] Configuration system
* [x] Structured logging
* [x] Unit testing
* [x] GitHub Actions
* [x] gRPC server
* [ ] PostgreSQL integration
* [ ] User registration
* [ ] User login
* [ ] JWT authentication
* [ ] Refresh tokens
* [ ] Email verification
* [ ] Docker support
* [ ] Jenkins pipeline

---

## License

This project is licensed under the MIT License.