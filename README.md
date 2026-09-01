# GhanaSchools

`GhanaSchools` is an independent Digital Ghana public-infrastructure registry. Its first release is a provenance-complete, GTEC-listed public-university subset with stable IDs and pinned GhanaGeo place references. Canonical web and API surfaces remain pre-release until production evidence supports a lifecycle transition.

The beta does not claim comprehensive school coverage, permanent accreditation, rankings or admissions authority.

## Before implementation

1. Record the problem, users, non-goals, source rights and acceptance evidence in `agent_plan.md`.
2. Replace the placeholder source-register record only after authority and licence review.
3. Add domain contracts before transport or UI code.
4. Keep deployments fail-closed until required provider values exist.

## Verification

Run `pnpm install` and `pnpm check` for governance validation, Go unit/API tests and vetting, dataset invariants, TypeScript checks, and the production web build.
