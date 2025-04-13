## 📁 Project Structure

This project follows an idiomatic Go folder structure inspired by real-world, production-grade REST APIs. It is designed to be modular, testable, and scalable while using only the Go standard library.

    social-media-api/
    ├── cmd/                # Application entrypoints (e.g., server startup code)
    │   └── api/            # Main API service executable
    │       └── main.go
    │
    ├── internal/           # Application-specific internal packages (not importable by other modules)
    │   ├── auth/           # Authorization logic (e.g., session/JWT management)
    │   ├── user/           # Domain logic for user management
    │   ├── post/           # Domain logic for social media posts
    │   └── middleware/     # Custom HTTP middleware (auth, logging, recovery, etc.)
    │
    ├── api/                # Versioned HTTP route handlers
    │   └── v1/             # Version 1 of the API
    │       ├── handlers.go
    │       ├── models.go
    │       └── routes.go
    │
    ├── pkg/                # Optional reusable public packages
    │   └── utils/          # Utility functions, hashing, validation, etc.
    │
    ├── docs/               # Project documentation, specifications, or diagrams
    ├── go.mod              # Go module definition
    ├── go.sum              # Dependency verification checksums
    └── README.md           # You're here!

> This structure helps separate concerns and promotes maintainability as the project grows. Business logic lives under `internal/`, while HTTP-layer logic and versioning are handled under `api/`.
