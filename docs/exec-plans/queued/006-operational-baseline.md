# Exec Plan 006: Operational Baseline

## Status

Queued.

## Branch

`feature/006-operational-baseline`

## Objective

Add local operations and minimal runtime visibility.

## Scope

- Add Docker Compose for PostgreSQL and Redis.
- Add `Makefile` targets.
- Add request ID middleware.
- Add structured logging.
- Add readiness endpoint.
- Add basic metrics endpoint.
- Update operations documentation.

## Out of scope

- Grafana dashboards.
- Alerting rules.
- Kubernetes manifests.

## Validation target

```bash
make validate
```

