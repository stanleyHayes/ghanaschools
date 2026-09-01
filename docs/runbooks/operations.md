# Operations baseline

No production service exists yet.

Before beta, define and verify:

- health, readiness and dependency checks;
- structured logs with request/correlation identifiers and secret redaction;
- latency, error, saturation and domain-correctness signals;
- alert routes with an acknowledged owner;
- backup/restore where state exists;
- least-privilege provider credentials and rotation;
- rate limits, abuse controls and audit trails for privileged publication;
- rollback and incident communication steps.

Deployment configuration is a skeleton until immutable build, smoke, TLS, rollback and provider evidence are recorded in `release-evidence.md`.
