# GhanaSchools execution ledger

Status: Public beta
Canonical hostname: `schools.digitalghana.dev`

## Product definition gate

- [x] Problem, users and non-goals approved in the supplied brief and `docs/product-definition.md`.
- [x] GTEC/GhanaGeo reference rights and publication boundaries recorded; no source document or bulk external dataset is redistributed.
- [x] Domain model, field provenance and deterministic fixtures implemented and passing.
- [x] Privacy, child-safety and misuse review complete; no personal, student, contact, admissions or inferred quality data is published.
- [x] Independent web/API boundary, constrained REST/GraphQL and TypeScript client scope approved.

## Live task board

| ID | Task | Status | Owner | Dependency | Evidence |
|---|---|---|---|---|---|
| P-0.1 | Product definition and source review | Done | Codex | — | GTEC public-university source, pinned GhanaGeo reference and owner brief recorded with explicit reference-only decisions |
| P-0.2 | Domain contracts and fixtures | Done | Codex | P-0.1 | 16-record provenance-complete dataset; stable-ID, alias, Geo-reference, duplicate and sensitive-field invariants pass |
| P-1.1 | Verified-subset implementation | Done — broader source classes remain gated | Codex | P-0.2 | Go REST/GraphQL, TypeScript client and Next.js directory pass; production Chrome proves alias search, custom region filter, fonts, safety copy, no overflow and zero prohibited native controls |
| P-2.1 | Production web/API release | Done — stable gates remain | Codex | P-1.1 | CI `33523973473`; Vercel `dpl_GqChhh8NFDVoC9d7LMBTGaXCnC7e`; Render `dep-dabep749v7es73cvajug`; canonical TLS/contract/SEO and rollback/restore evidence recorded |
