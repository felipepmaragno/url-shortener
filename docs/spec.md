# Behavior Spec

This document owns externally observable behavior.

## Create short URL

```http
POST /api/v1/shorten
Content-Type: application/json

{
  "url": "https://example.com/path",
  "custom_alias": "optional-alias",
  "expires_at": "2026-12-31T00:00:00Z"
}
```

`custom_alias` and `expires_at` are planned features. The first MVP may omit them.

Success:

```http
201 Created
```

```json
{
  "code": "b7Xk91",
  "short_url": "https://localhost:8080/b7Xk91",
  "long_url": "https://example.com/path"
}
```

Errors:

- `400 Bad Request`: invalid URL or invalid alias.
- `409 Conflict`: custom alias already exists.
- `429 Too Many Requests`: rate limited.
- `500 Internal Server Error`: unexpected failure.

## Redirect

```http
GET /{code}
```

Success:

```http
302 Found
Location: https://example.com/path
```

Errors:

- `404 Not Found`: unknown code.
- `410 Gone`: expired or disabled code, if expiration/disable support is enabled.

## Invariants

- One short code must never point to two different long URLs.
- Generated public codes must use base62.
- Generated public codes must have minimum length 5.
- Codes with length 1–4 are reserved for internal, premium, or manual aliases.
- Cache is not source of truth.
- Primary storage is source of truth.
