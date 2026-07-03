# Product Scope

## Purpose

Build a URL shortener service that creates compact short URLs and redirects them to their original URLs.

## Primary user

A user or API client that wants to convert a long URL into a short URL and later redirect through that short URL.

## Core promise

Given a valid long URL, the service creates a compact short code. Given an active short code, the service redirects to the original URL.

## In scope

- Create short URLs.
- Redirect short URLs.
- Validate submitted URLs.
- Generate compact unique short codes.
- Persist mappings.
- Cache hot redirects.
- Support expiration and custom aliases as product features.
- Provide basic operational endpoints and metrics.

## Out of scope for the first implementation

- Full user account system.
- Frontend dashboard.
- Full analytics product.
- Malware scanning implementation.
- Multi-region deployment.
- Kubernetes deployment.
