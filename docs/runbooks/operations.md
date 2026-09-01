# Operations baseline

The public directory runs on Vercel at `schools.digitalghana.dev`. The stateless Go API runs on Render at `api-schools.digitalghana.dev`; the provider hostname is `ghanaschools-api.onrender.com`. There is no application database, privileged interface or user-generated content in this beta.

## Routine checks

- Confirm web, robots, sitemap, favicon, manifest and Open Graph image return 200.
- Confirm `/health` reports 16 records and dataset `2026.09.01-beta.1`.
- Search `KNUST`, filter `region=Central`, retrieve `unimac`, and compare constrained GraphQL parity.
- Confirm CORS allows only `https://schools.digitalghana.dev`.
- Recheck the GTEC source before dataset changes and visibly retain the reviewed-at date.
- Run `pnpm check` and require GitHub Quality success before release.

## Rollback

Use immutable Vercel deployment IDs and Render deploy IDs. After rollback, smoke both canonical hostnames and restore the intended release. A custom Vercel alias may need explicit reassignment. Render rollback uses `POST /v1/services/{serviceId}/rollback` with a known-good `deployId`.

Before a stateful admin, correction, school-claim or broader import system, define and verify:

- health, readiness and dependency checks;
- structured logs with request/correlation identifiers and secret redaction;
- latency, error, saturation and domain-correctness signals;
- alert routes with an acknowledged owner;
- backup/restore where state exists;
- least-privilege provider credentials and rotation;
- rate limits, abuse controls and audit trails for privileged publication;
- rollback and incident communication steps.
