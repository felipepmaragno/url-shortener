# Architecture

This document owns runtime structure and component responsibilities.

## Target architecture

```text
Client
  -> DNS/CDN
  -> Load Balancer
  -> Go API service
  -> Redis redirect cache
  -> PostgreSQL primary storage
```

ID generation:

```text
Go API service
  -> Redis INCRBY range allocator
  -> local ID range
  -> base62 encoder
```

## Component responsibilities

### Go API service

- Validate create requests.
- Generate short codes.
- Persist URL mappings.
- Resolve codes for redirects.
- Apply cache-aside lookup.
- Return HTTP responses.

### PostgreSQL

- Source of truth for URL mappings.
- Enforce unique `short_code`.
- Store expiration and disabled state when implemented.

### Redis redirect cache

- Cache hot `code -> long_url` mappings.
- Disposable; fallback to PostgreSQL on miss or failure.

### Redis ID range allocator

- Allocate numeric ID ranges with `INCRBY`.
- Correctness-critical.
- Should be logically or operationally separate from redirect cache.

## Key design choices

- Use `302` redirects by default.
- Use variable-length base62 codes with minimum generated length 5.
- Reserve 1–4 character codes.
- Use Redis `INCRBY` range allocation for production-minded ID generation.
- Keep analytics and fraud detection outside the core implementation.

