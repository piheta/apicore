# apicore

[![Go Report Card](https://goreportcard.com/badge/github.com/piheta/apicore)](https://goreportcard.com/report/github.com/piheta/apicore)
[![Go Test](https://github.com/piheta/apicore/actions/workflows/test.yml/badge.svg)](https://github.com/piheta/apicore/actions/workflows/test.yml)
[![Go Lint](https://github.com/piheta/apicore/actions/workflows/lint.yml/badge.svg)](https://github.com/piheta/apicore/actions/workflows/lint.yml)
[![CodeQL](https://github.com/piheta/apicore/actions/workflows/codeql.yml/badge.svg)](https://github.com/piheta/apicore/security/code-scanning)

A lightweight Go toolkit for building APIs. Provides utilities for structured error handling, error mapping, middleware support, and structured request logging.

## Features

- **Error handling** - Structured API errors with metadata tracing and propagation.
- **Logging** - Automatic request logging with error details and metadata extraction.
- **Middleware support** - Chainable middleware for request processing
- **Response formatting** - Standardized API response structure

## Example Usage
```go
package main

import (
	"net/http"

	"github.com/piheta/apicore/apierr"
	"github.com/piheta/apicore/metaerr"
	"github.com/piheta/apicore/middleware"
	"github.com/piheta/apicore/response"
)

func main() {
	mux := http.NewServeMux()

	// Register handlers using public middleware
	mux.Handle("GET /api/ping", middleware.Public(Ping))
	mux.Handle("GET /api/err", middleware.Public(Err))

	server := &http.Server{
		Addr:         ":8080",
		Handler:      middleware.RequestLogger(mux), // Wrap handlers with logging middleware
	}

	server.ListenAndServe()
}

// Working handler
func Ping(w http.ResponseWriter, r *http.Request) error {
	return response.JSON(w, 200, "pong")
}

// Failing handler
func Err(w http.ResponseWriter, r *http.Request) error {
	err := apierr.NewError(404, "not_found", "user not found") // APIError (code, type, msg)

	// MetaErr, wraps error with additional key-value pair metadata for logging
	return metaerr.WithMetadata(err, "user_id", "123", "email", "user@example.com")
}
```

### Logging
```
2025/11/28 22:25:18 INFO REQ status=200 ms=0.09 ip=[::1]:51420 method=GET path=/api/ping
2025/11/28 22:25:24 WARN REQ status=404 ms=0.16 ip=[::1]:51425 method=GET path=/api/err error_detail="user not found" user_id=123 email=user@example.com error="Not Found"
```

## Status
⚠️ **Pre-release**: This library is under active development.
Breaking changes may occur before v1.0.0 is released.
Pin to a specific version when using as a dependency.
