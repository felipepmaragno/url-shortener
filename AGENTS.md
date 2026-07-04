# Agent Guide

This file tells future maintainers and AI agents how to work safely in this repository.

## Project classification

This is a service project. Keep the implementation focused and use documentation to preserve current behavior, design decisions, tradeoffs, and validation evidence.

## Documentation ownership

- `README.md`: quick project orientation and links.
- `PROGRESS.md`: current verified state and next safe starting point.
- `docs/product.md`: purpose, users, scope, and non-goals.
- `docs/spec.md`: externally observable behavior and invariants.
- `docs/architecture.md`: runtime structure and component responsibilities.
- `docs/engineering.md`: Go implementation standards, test strategy, and quality gates.
- `docs/operations.md`: local run, validation, and failure-handling notes.
- `docs/limitations.md`: accepted limitations and future work boundaries.
- `docs/adr/`: durable decision history.
- `docs/exec-plans/`: bounded implementation plans.
- `docs/learnings/`: lessons and implementation notes that are not current authority.
- `docs/spikes/`: investigations not yet accepted as design.

Avoid duplicating durable facts. Link to the owning document instead.

## Current implementation rule

Implementation has not started. Before adding code:

1. Update or create an execution plan under `docs/exec-plans/active/`.
2. Keep the increment small and reviewable.
3. Add tests with the feature.
4. Update `PROGRESS.md` when verified state changes.

## Branch and PR workflow

- `main` is the integration branch.
- Each increment branch starts from `main`.
- Each increment is merged back to `main`.
- Branch names use `feature/...` for feature work and `fix/...` for corrections, for example `feature/001-simple-mvp`.
- End each increment with a pushed branch and a draft pull request.

## Expected engineering style

- Keep HTTP handlers thin.
- Put business behavior in application services.
- Hide persistence behind interfaces.
- Use `context.Context` for request-scoped work.
- Start with the Go standard library before adding third-party libraries.
- Follow TDD for each increment.
- Keep `make validate` as the local quality gate once available.
- Treat PostgreSQL as source of truth.
- Treat redirect cache as disposable.
- Treat ID allocation as correctness-critical.
