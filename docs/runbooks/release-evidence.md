# GhanaSchools public beta evidence — 2026-09-01

## Release identity

- Repository: `https://github.com/stanleyHayes/ghanaschools`
- Application commit: `dbddac97f1975753ab3e412a89bb71c5359d572c`
- Dataset: `2026.09.01-beta.1`, reviewed 2026-09-01
- GhanaGeo reference dataset: `2026.08.3-ulid`, pinned to GeoGhana commit `324916ee6e5221a183095348962ff6c7c2e29d98`
- GitHub Actions Quality run: `33523973473`, success

## Provider inventory

- Vercel project: `hayfordstanleys-projects/ghanaschools` (`prj_TCepVcfcQrASONP57NzBv9kFmFO3`)
- Current web deployment: `dpl_GqChhh8NFDVoC9d7LMBTGaXCnC7e`
- Canonical web: `https://schools.digitalghana.dev`
- Render service: `srv-dabelrss728c73ag0o4g`
- Current API restore deploy: `dep-dabep749v7es73cvajug`
- Render custom domain: `cdm-dabem4n10e5c73aohb80`, verified
- Canonical API: `https://api-schools.digitalghana.dev`

## Dataset and contract verification

- Sixteen public-university records have unique stable IDs and field-level GTEC/GhanaGeo provenance.
- Tests prove alias behavior, duplicate-name preservation, pinned GhanaGeo region resolution and absence of sensitive/inferred fields.
- REST `KNUST` search returned only `knust`; Central filter returned `ucc` and `uew`; `unimac` retained four documented aliases.
- Constrained GraphQL returned the same `knust` record and dataset version as REST.
- API CORS allows `https://schools.digitalghana.dev`.

## Browser, accessibility and SEO verification

- Production Chrome alias search returned one KNUST record; the custom Radix region picker exposed eleven choices and Greater Accra returned five records.
- Outfit body/UI and Geist Mono identifiers are active; Newsreader is the editorial title face.
- No native select, dialog, checkbox, radio, date or time control renders. No horizontal overflow was observed.
- Canonical, Open Graph, Twitter and JSON-LD metadata are present.
- Web, robots, sitemap, SVG favicon and web manifest returned 200 with expected types; the Open Graph PNG is 1200×630.

## Rollback evidence

- Vercel rolled back from `dpl_GqChhh8NFDVoC9d7LMBTGaXCnC7e` to `dpl_5eJsBrcS6AvvTSSzgrp3Lv5Nk3bL`; canonical smoke passed; the current deployment alias was restored and smoke passed.
- Render rolled back to original deploy `dep-dabelsks728c73ag0q4g`, producing rollback deploy `dep-dabeoth42hec73aos9hg`; canonical health passed.
- Render restored the intended release line from `dep-dabeogc9v7es73cv4av0`, producing `dep-dabep749v7es73cvajug`; canonical health passed.

## Known limits

- Coverage is limited to the reviewed GTEC public-university list; it is not a comprehensive Ghana school directory.
- A source listing is not a permanent accreditation guarantee. Users must recheck time-sensitive decisions with GTEC.
- GCTU and GIMPA use the broader Accra GhanaGeo reference with a visible location note; exact Tesano/Achimota references remain review items.
- No programmes, coordinates, contacts, fees, rankings, admissions data, student data or inferred quality fields are published.
- Broader pre-tertiary/private/technical/college sources, nearby search, gRPC, hooks and correction/admin workflows remain gated.
