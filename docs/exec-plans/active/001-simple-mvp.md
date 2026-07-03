# Exec Plan 001: Simple MVP

## Status

Active.

## Branch

`increment/001-simple-mvp`

## Objective

Build the smallest runnable URL shortener with in-memory storage.

## Scope

- Create Go module.
- Add base62 encoder.
- Add minimum generated code length of 5.
- Add in-memory URL repository.
- Add application service for:
  - creating a short URL;
  - resolving a short code.
- Add HTTP endpoints:
  - `POST /api/v1/shorten`;
  - `GET /{code}`;
  - `GET /healthz`.
- Add tests for encoder, service, and HTTP handlers.

## Out of scope

- PostgreSQL.
- Redis.
- Docker Compose.
- Custom aliases.
- Expiration.
- Metrics.
- Authentication.

## Validation target

```bash
go test ./...
```

## Completion criteria

- The service can create and resolve short URLs in memory.
- Unknown codes return `404`.
- Redirects use `302` and the `Location` header.
- Tests pass.

