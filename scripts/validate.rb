#!/usr/bin/env ruby

require "json"
require "pathname"

root = Pathname.new(__dir__).parent
required = %w[README.md AGENTS.md agent_plan.md SECURITY.md LICENSE package.json pnpm-lock.yaml tsconfig.json next.config.ts vercel.json go.mod render.yaml data/schools.json internal/catalog/catalog.go cmd/api/main.go sdk/typescript/index.ts tests/dataset.test.mjs app/layout.tsx app/page.tsx app/directory.tsx app/styles.css app/icon.svg app/opengraph-image.tsx contracts/README.md contracts/openapi.yaml contracts/graphql/schema.graphql docs/product-definition.md docs/adr/0001-product-boundary.md docs/governance/source-register.json docs/runbooks/operations.md docs/runbooks/release-evidence.md infra/vercel.json]
missing = required.reject { |path| root.join(path).file? }
abort "missing required files: #{missing.join(', ')}" unless missing.empty?

source_register = JSON.parse(root.join("docs/governance/source-register.json").read)
abort "source register must contain at least one record" if source_register.fetch("sources", []).empty?
abort "source register contains blocked or unknown decisions" if source_register.fetch("sources").any? { |source| ["blocked", "unknown"].include?(source.fetch("publicationDecision")) }

dataset = JSON.parse(root.join("data/schools.json").read)
abort "beta must contain exactly 16 records" unless dataset.fetch("schools").length == 16
abort "dataset version missing" if dataset.fetch("datasetVersion", "").empty?

contents = required.map { |path| root.join(path).read }.join("\n")
template_marker = ["__", "PRODUCT_"].join
abort "unresolved template token" if contents.include?(template_marker)
private_key_pattern = Regexp.new(["-----BEGIN", ".*PRIVATE KEY-----"].join(" "))
abort "possible private key" if contents.match?(private_key_pattern)

puts "GhanaSchools validation passed: #{dataset.fetch('schools').length} records"
