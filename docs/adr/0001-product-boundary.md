# ADR-0001: Independent product boundary

Status: Accepted

## Decision

GhanaSchools owns its repository, release lifecycle, data, credentials, contracts and operational evidence. It may consume another Digital Ghana product only through a versioned public contract or pinned dataset artifact.

The canonical web hostname is `schools.digitalghana.dev`. Additional API or operational hostnames require an explicit ADR and portfolio registry entry before deployment.

The first release is a deliberately narrow GTEC-listed public-university subset. It is a discovery and interoperability registry, not an admissions, accreditation-decision, ranking, safeguarding, student-record or school-claim system. The canonical API hostname is `api-schools.digitalghana.dev`, recorded in the central portfolio registry.

Every published field names its source. Location linkage uses a pinned GhanaGeo dataset contract; GhanaSchools never reads GhanaGeo's database or invents campus coordinates from a town name.

## Consequences

Failures remain isolated, histories remain understandable, and a portfolio-wide platform outage is not created by convenience. Some configuration and small primitives may be repeated until two proven consumers justify a versioned shared package.
