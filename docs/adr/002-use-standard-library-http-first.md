# ADR 002: Use The Go Standard Library HTTP Stack First

## Status

Accepted.

## Context

The service initially needs a small HTTP surface:

- `POST /api/v1/shorten`;
- `GET /{code}`;
- `GET /healthz`.

The initial route set is simple. The project should avoid unnecessary dependencies until routing, middleware, or request handling requirements justify them.

## Decision

Use the Go standard library HTTP stack first:

- `net/http`;
- `http.ServeMux`;
- `http.Handler`;
- `http.HandlerFunc`.

Third-party routers or frameworks may be introduced later through a separate ADR if the standard library becomes a source of complexity.

## Rationale

The standard library is enough for the MVP because:

- the route set is small;
- `net/http` is stable and widely understood;
- handlers and middleware can be tested directly;
- avoiding dependencies keeps the first implementation easier to review;
- future migration to a router remains straightforward if requirements grow.

## Alternatives considered

### `chi`

`chi` is idiomatic, lightweight, and a good option for larger route sets or richer middleware composition. It is not needed for the initial MVP.

### Full web framework

Frameworks can speed up some projects but add conventions and dependencies that are unnecessary for this service at the current scope.

## Consequences

- Initial handlers should use standard-library request parsing and response writing.
- Middleware should be implemented using `func(http.Handler) http.Handler`.
- Route complexity should be monitored as the API grows.
- Adding a third-party router later requires a focused ADR.

