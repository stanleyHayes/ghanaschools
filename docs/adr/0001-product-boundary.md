# ADR-0001: Independent product boundary

Status: Accepted

## Decision

GhanaSchools owns its repository, release lifecycle, data, credentials, contracts and operational evidence. It may consume another Digital Ghana product only through a versioned public contract or pinned dataset artifact.

The canonical web hostname is `schools.digitalghana.dev`. Additional API or operational hostnames require an explicit ADR and portfolio registry entry before deployment.

## Consequences

Failures remain isolated, histories remain understandable, and a portfolio-wide platform outage is not created by convenience. Some configuration and small primitives may be repeated until two proven consumers justify a versioned shared package.
