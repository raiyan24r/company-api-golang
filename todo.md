# Company API — Technical TODO

This checklist tracks what’s already in place and what remains to make the service production-ready per requirements.

## Done
- [x] Project modules and dependencies (`go.mod`, vendored deps)
- [x] Base project structure (`app/api`, handlers, routes, requests/responses)
- [x] Company repository and DB layer scaffolding (`business/database/companyrepo.go`)
- [x] DB migration scaffold for companies table (`business/database/migration/000001_create_companies_table.*.sql`)
- [x] Configuration files present (`config.yml`, `config.yml.dist`) using `viper`
- [x] Structured logging foundation (`foundation/logger` using `zap`)
- [x] Middleware scaffolding including request ID (`foundation/middleware/requestid.go`)
- [x] Docker assets present (`Dockerfile`, `docker-compose.yml`)
- [x] Makefile with common targets (`Makefile`)
- [x] Basic database utilities and a unit test scaffold (`foundation/database/db.go`, `db_test.go`)
- [x] README baseline (`README.md`)

## To Do — Core API
- [ ] Define canonical `Company` domain model with validation:
  - `id` (UUID, required)
  - `name` (string ≤15 chars, required, unique)
  - `description` (string ≤3000 chars, optional)
  - `employees` (int, required)
  - `registered` (bool, required)
  - `type` (enum: Corporations | NonProfit | Cooperative | Sole Proprietorship)
- [ ] REST endpoints (Chi router) with JSON:
  - [ ] POST `/companies` — Create
  - [ ] PATCH `/companies/{id}` — Partial update
  - [ ] DELETE `/companies/{id}` — Delete
  - [ ] GET `/companies/{id}` — Get one
- [ ] Request/response DTOs and mappers between DTO ↔ domain ↔ storage
- [ ] Input validation: length checks, enum validation, required fields, type safety
- [ ] Error model: consistent error responses (`app/api/handler/response/error.go`)

## To Do — Authentication/Authorization
- [ ] JWT authentication middleware (mutations require auth):
  - [ ] Verify tokens (HS256 using configured secret, or RS256 using keypair)
  - [ ] Extract subject/claims into request context
  - [ ] Protect `POST`, `PATCH`, `DELETE` routes; allow anonymous `GET`
- [ ] Config entries for JWT (`issuer`, `audience`, `alg`, `secret/publicKey`)
- [ ] Optional: role-based authorization checks for future extension

## To Do — Persistence
- [ ] Finalize companies table schema with constraints:
  - [ ] `name` UNIQUE index
  - [ ] `type` constrained to enum values (DB CHECK or app-level guard)
  - [ ] `id` UUID primary key (stored as `CHAR(36)` or `BINARY(16)`)
- [ ] Implement repository methods:
  - [ ] `CreateCompany(ctx, Company)`
  - [ ] `PatchCompany(ctx, id, PatchCompany)` with partial update logic
  - [ ] `DeleteCompany(ctx, id)`
  - [ ] `GetCompany(ctx, id)`
- [ ] Transactions where appropriate; proper `context.Context` usage
- [ ] Migrate on startup or provide `make migrate` target

## To Do — Events (Kafka)
- [ ] Produce events on mutating operations:
  - [ ] `CompanyCreated`
  - [ ] `CompanyUpdated`
  - [ ] `CompanyDeleted`
- [ ] Define event schemas (JSON) and topics
- [ ] Kafka producer setup with retries and backoff
- [ ] Add Kafka/Zookeeper services in `docker-compose.yml`
- [ ] Config entries for Kafka brokers and topic names

## To Do — Infrastructure & Ops
- [ ] Docker: production-ready image (multi-stage build, minimal runtime layer)
- [ ] Docker Compose: services for API, DB (e.g., MySQL/Postgres), Kafka
- [ ] Health endpoint `/health` and readiness checks
- [ ] Graceful shutdown handling (HTTP server, DB, Kafka producer)
- [ ] Config management via `viper` + environment variables (`gotenv` support)
- [ ] Observability: structured logs with correlation IDs
- [ ] Optional: metrics (Prometheus) and tracing (OTel) for bonus points

## To Do — Testing & Quality
- [ ] Unit tests for handlers, services, repositories
- [ ] Integration tests (API + DB + Kafka) via docker-compose or Testcontainers
- [ ] Linter: add `golangci-lint` with config `.golangci.yml`
- [ ] CI: GitHub Actions to run build, lint, tests on PRs
- [ ] Security checks: vet, staticcheck (as available)

## To Do — Documentation
- [ ] Update `README.md` with clear setup/run instructions:
  - [ ] Local run (make, env, config)
  - [ ] Docker Compose (DB, Kafka)
  - [ ] Production image build
  - [ ] Authentication usage (JWT examples)
  - [ ] API examples (curl)
- [ ] Optional: OpenAPI spec for endpoints

## To Do — Makefile & Scripts
- [ ] Make targets: `build`, `run`, `test`, `lint`, `format`, `docker-build`, `compose-up`, `compose-down`, `migrate`
- [ ] Pre-commit hooks (format/lint)

## Implementation Notes
- Use Chi for routing, Zap for logging, Viper for config
- Prefer context-aware DB access and timeouts
- Validate strictly at the edge (handler layer); keep domain clean
- Emit events after successful DB commits to avoid outbox inconsistency; consider outbox pattern if needed
- Keep public APIs stable and avoid unnecessary renames
