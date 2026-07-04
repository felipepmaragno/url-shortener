# Exec Plan 002: PostgreSQL Persistence

## Status

Queued.

## Branch

`feature/002-postgres-persistence`

## Objective

Replace in-memory storage with durable PostgreSQL persistence.

## Scope

- Add PostgreSQL dependency.
- Add migrations.
- Add `urls` table.
- Add PostgreSQL repository.
- Enforce unique `short_code`.
- Add repository integration tests.

## Out of scope

- Redis cache.
- Redis ID allocation.
- Custom aliases.
- Expiration behavior.

## Validation target

```bash
go test ./...
```

