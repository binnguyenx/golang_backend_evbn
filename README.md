# golang_backend_evbn

Golang backend project.

## Project Structure

```
.
├── cmd/                # Application entrypoints
├── internal/
│   ├── handler/        # HTTP handlers
│   ├── model/          # Data models
│   ├── repository/     # Data access layer
│   └── service/        # Business logic
├── go.mod
└── README.md
```

## Getting Started

```bash
go mod tidy
go run ./cmd/...
```
