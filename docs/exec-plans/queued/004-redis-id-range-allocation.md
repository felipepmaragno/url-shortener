# Exec Plan 004: Redis ID Range Allocation

## Status

Queued.

## Branch

`increment/004-redis-id-range-allocation`

## Objective

Replace simple local ID generation with Redis `INCRBY` range allocation.

## Scope

- Add Redis dependency.
- Add ID allocator interface.
- Implement Redis `INCRBY` allocator.
- Consume allocated ID ranges locally.
- Continue generating IDs while a local range exists if Redis is temporarily unavailable.
- Add allocator tests.

## Out of scope

- Redirect caching.
- Rate limiting.
- Multi-region ID allocation.

## Validation target

```bash
go test ./...
```

