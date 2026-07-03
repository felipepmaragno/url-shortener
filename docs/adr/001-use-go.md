# ADR 001: Use Go For The Service

## Status

Accepted.

## Context

The service needs to expose HTTP endpoints, perform simple request validation, access storage/cache dependencies, and remain easy to test and operate. The implementation should keep concurrency, cancellation, timeouts, and dependency boundaries explicit.

The expected service shape is:

- HTTP API for creating and resolving short URLs.
- Application layer for URL creation and redirect resolution.
- Repository and cache interfaces.
- PostgreSQL and Redis integrations in later increments.
- Operational endpoints and metrics.

## Decision

Use Go as the implementation language.

## Rationale

Go fits this service because:

- the standard library has strong HTTP support;
- explicit error handling keeps failure paths visible;
- `context.Context` is a good fit for request-scoped cancellation and timeouts;
- interfaces make storage, cache, and ID allocation dependencies easy to test;
- compiled binaries are simple to deploy;
- goroutines and channels are available if asynchronous work is added later;
- Go projects can stay small without requiring a heavy framework.

## Alternatives considered

### Node.js / TypeScript

Strong ecosystem and fast development. Rejected for this project because the service benefits from Go's simple deployment model, explicit concurrency primitives, and standard-library HTTP capabilities.

### Java / Kotlin

Mature service ecosystem and strong operational tooling. Rejected for the initial implementation because it adds more framework and runtime weight than this service needs.

### Python

Fast to prototype and easy to read. Rejected for the primary implementation because the target service benefits from static typing, simple binary deployment, and Go's HTTP/concurrency model.

## Consequences

- The project will use idiomatic Go package boundaries.
- The service should avoid unnecessary framework dependencies until there is a concrete need.
- Tests should cover application behavior independently from HTTP and infrastructure adapters.
- Future ADRs should cover storage, cache, ID allocation, redirect semantics, and observability separately.

