# GhanaSchools product definition

## Problem and users

Education applications repeatedly rebuild institution lists without stable identifiers, source dates or honest coverage labels. GhanaSchools gives developers, researchers and civic-technology teams a provenance-complete registry contract that can grow through reviewed releases.

## Public beta scope

The first release contains only the public universities listed in the reviewed GTEC current-institutions source. Each record has a GhanaSchools stable ID, name, aliases where present in the source, ownership/type, source-backed locality text and a pinned GhanaGeo place reference.

Search supports names and aliases; filters support location and institution type. REST, constrained GraphQL and a dependency-free TypeScript client expose the same dataset version.

## Non-goals and safety

- No claim that the subset covers every school or tertiary institution in Ghana.
- No student, parent, staff, child, contact-person or admissions data.
- No rankings, fees, programmes, results, reviews or inferred quality claims.
- No invented campus coordinates, phone numbers, emails or codes.
- “Listed by GTEC source” is not rendered as a permanent accreditation guarantee; users are directed to recheck time-sensitive status with GTEC.
- No institution claim, edit or verification workflow before authentication, audit and safeguarding review.

## Stable identity and provenance

IDs are product-owned slugs and do not change when a display name changes. Aliases are additive. Every field lists one or more source IDs, and every release names the source review date plus GhanaGeo dataset version. Conflicting or duplicate-name records remain distinct and enter review; they are never merged by name alone.

## Definition-gate acceptance

- Every record and field has machine-readable provenance.
- All GhanaGeo references resolve in the pinned dataset and select the intended region/locality among ambiguous names.
- Search finds documented aliases.
- Duplicate names are permitted by schema and IDs remain unique.
- No coordinate or accreditation-status inference is present.
