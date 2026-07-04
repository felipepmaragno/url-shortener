# Exec Plan 005: Redis Redirect Cache

## Status

Queued.

## Branch

`feature/005-redis-redirect-cache`

## Objective

Add Redis cache-aside behavior for redirect resolution.

## Scope

- Add redirect cache interface.
- Implement Redis cache for `code -> long_url`.
- Check cache before PostgreSQL on redirect.
- Populate cache after PostgreSQL hit.
- Fall back to PostgreSQL if Redis cache fails.
- Add cache behavior tests.

## Out of scope

- CDN caching.
- Analytics events.
- Cache warming jobs.

## Validation target

```bash
go test ./...
```

