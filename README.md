# URL Shortener

A production-minded URL shortener service implemented in Go.

The project focuses on the core system design first:

- create a short URL from a long URL;
- redirect a short code to the original URL;
- generate compact unique codes;
- store mappings durably;
- use caching for the read-heavy redirect path;
- document tradeoffs and operational behavior clearly.

## Current status

The simple in-memory MVP is implemented on the active feature branch.

Start here:

- [Product scope](docs/product.md)
- [Behavior spec](docs/spec.md)
- [Architecture](docs/architecture.md)
- [Engineering standards](docs/engineering.md)
- [Operations](docs/operations.md)
- [Progress](PROGRESS.md)

Supporting design references:

- [Current Draw.io diagram](system-design.drawio)
- [Full single-page diagram](system-design-full-single-page.drawio)

## Planned implementation stack

- Go HTTP service
- PostgreSQL as source of truth
- Redis for redirect cache
- Redis `INCRBY` range allocation for generated IDs
- Docker Compose for local dependencies

## Planned validation

Local validation:

```bash
make test
make test-race
make run
make validate
```
