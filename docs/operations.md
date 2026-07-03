# Operations

This document owns local run, validation, and operational notes.

## Current state

Implementation has not started. Commands below are planned targets.

## Planned local dependencies

- PostgreSQL
- Redis

## Planned commands

```bash
make up          # start dependencies
make migrate     # run database migrations
make run         # run service locally
make test        # run tests
make validate    # run full local validation
```

## Planned smoke test

```bash
curl -X POST http://localhost:8080/api/v1/shorten \
  -H 'Content-Type: application/json' \
  -d '{"url":"https://example.com"}'

curl -i http://localhost:8080/{code}
```

## Failure behavior targets

- Redis cache down: redirect falls back to PostgreSQL.
- Redis allocator down: creates continue while local ID range remains; fail safely after range exhaustion.
- PostgreSQL down: creates fail; redirects work only for cached active links.

