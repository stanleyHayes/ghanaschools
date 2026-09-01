# GhanaSchools execution ledger

Status: Planning  
Canonical hostname: `schools.digitalghana.dev`

## Product definition gate

- [ ] Problem, users and non-goals approved.
- [ ] Source authority and licence review complete.
- [ ] Domain model and deterministic acceptance fixtures approved.
- [ ] Security, privacy and misuse risks reviewed.
- [ ] API/UI scope and deployment boundary approved.

## Live task board

| ID | Task | Status | Owner | Dependency | Evidence |
|---|---|---|---|---|---|
| P-0.1 | Product definition and source review | Ready | Unassigned | — | Product-specific source record required |
| P-0.2 | Domain contracts and fixtures | Blocked | Unassigned | P-0.1 | No implementation before definition gate |
| P-1.1 | Implementation | Blocked | Unassigned | P-0.2 | Tests must prove domain behavior |
| P-2.1 | Production release | Blocked | Unassigned | P-1.1 | Smoke, security, rollback and operations evidence required |
