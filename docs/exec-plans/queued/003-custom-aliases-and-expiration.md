# Exec Plan 003: Custom Aliases And Expiration

## Status

Queued.

## Branch

`feature/003-custom-aliases-and-expiration`

## Objective

Add product behavior for custom aliases and expiring links.

## Scope

- Accept optional `custom_alias` in create request.
- Validate alias format.
- Reserve system aliases.
- Return `409 Conflict` when alias already exists.
- Accept optional `expires_at`.
- Return `410 Gone` for expired links.
- Add tests for alias and expiration behavior.

## Out of scope

- Authentication and ownership.
- Editing existing URLs.
- Analytics.

## Validation target

```bash
go test ./...
```

