# Engineering Standards

This document owns implementation standards, test strategy, and quality gates.

## Go implementation principles

- Prefer the Go standard library before adding third-party dependencies.
- Keep packages small and behavior-focused.
- Keep HTTP handlers thin.
- Put business behavior in application services.
- Keep persistence, cache, and ID allocation behind interfaces.
- Use `context.Context` for request-scoped work.
- Return explicit errors and map them to HTTP responses at the boundary.
- Avoid generic packages such as `utils`, `common`, or `helpers`.

The intended dependency direction is:

```text
cmd/urlshortener
  -> internal/httpapi
  -> internal/app
  -> internal/domain

internal/app
  -> repository/cache/id-generator interfaces

infrastructure adapters
  -> implement interfaces
```

## Initial package structure

The first implementation should start with:

```text
cmd/urlshortener/
internal/config/
internal/httpapi/
internal/app/
internal/domain/
internal/codegen/
internal/store/memory/
```

Later increments may add:

```text
internal/store/postgres/
internal/cache/redis/
internal/idalloc/redis/
internal/observability/
```

## Error handling

Domain and application errors should be stable enough for HTTP mapping.

Examples:

```text
invalid_url
not_found
internal_error
```

HTTP error responses use this envelope:

```json
{
  "error": {
    "code": "invalid_url",
    "message": "invalid URL"
  }
}
```

Internal errors may be wrapped with context, but HTTP responses must not leak implementation details.

## TDD workflow

Each increment should follow this loop:

1. Write a failing test for the next behavior.
2. Implement the smallest code that passes.
3. Refactor while tests stay green.
4. Add boundary tests for HTTP or infrastructure behavior when relevant.
5. Update owning docs when durable behavior changes.

For the MVP, the recommended test order is:

1. base62 encoder tests;
2. code generation/minimum-length tests;
3. URL validation tests;
4. application service tests;
5. HTTP handler tests;
6. config tests.

## Test layers

### Unit tests

Fast tests with no external dependencies.

Use for:

- base62 encoding;
- URL validation;
- application behavior;
- HTTP error mapping;
- in-memory repository behavior.

Command:

```bash
go test ./...
```

### Race tests

Use for code that may be used concurrently.

Command:

```bash
go test -race ./...
```

### Integration tests

Use real external dependencies such as PostgreSQL and Redis.

These should start when those dependencies are introduced.

Recommended build tag:

```go
//go:build integration
```

Command:

```bash
go test -tags=integration ./...
```

### End-to-end tests

Run the service and exercise HTTP behavior through the network boundary.

These should start once Docker Compose or an equivalent local runtime exists.

Command target:

```bash
make test-e2e
```

### Smoke tests

Small checks against a running service.

Examples:

- `GET /healthz`;
- create a URL;
- follow a redirect response and verify `Location`.

Command target:

```bash
make smoke
```

### Load tests

Load tests are not part of the first MVP. Add them after the redirect path, cache, and persistence behavior are stable.

Candidate tools:

- `k6`;
- `vegeta`;
- `hey`.

Load tests should be manual or scheduled, not required on every pull request.

## Quality gates

The local quality gate should be:

```bash
make validate
```

Initial target:

```text
format check
go vet ./...
go test ./...
go test -race ./...
```

Later targets:

```text
integration tests
e2e tests
smoke tests
coverage report
static analysis
```

## Coverage expectations

Coverage should prove behavior. Do not chase 100%.

Expected emphasis:

- `internal/domain`, `internal/codegen`, and `internal/app`: high unit coverage;
- `internal/httpapi`: behavior-focused handler coverage;
- infrastructure adapters: integration coverage;
- `cmd/urlshortener`: minimal coverage; validate through smoke/e2e tests.

## CI progression

### CI v1: MVP

Run on pull requests:

```text
format check
go vet ./...
go test ./...
go test -race ./...
```

### CI v2: PostgreSQL and Redis

Add service dependencies and run:

```text
unit tests
integration tests
```

### CI v3: runnable service

Add:

```text
build binary
start service
run smoke/e2e tests
```

### CI v4: performance validation

Add manual or scheduled load tests.

## Third-party dependency rule

Add a dependency only when it removes meaningful complexity or provides behavior that would be risky to maintain locally.

Before adding a dependency, document:

- what problem it solves;
- why the standard library is insufficient;
- maintenance and operational impact.

