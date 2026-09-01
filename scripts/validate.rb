#!/usr/bin/env ruby

require "json"
require "pathname"

root = Pathname.new(__dir__).parent
required = %w[README.md AGENTS.md agent_plan.md SECURITY.md LICENSE contracts/README.md docs/adr/0001-product-boundary.md docs/governance/source-register.json docs/runbooks/operations.md docs/runbooks/release-evidence.md infra/vercel.json]
missing = required.reject { |path| root.join(path).file? }
abort "missing required files: #{missing.join(', ')}" unless missing.empty?

source_register = JSON.parse(root.join("docs/governance/source-register.json").read)
abort "source register must contain at least one record" if source_register.fetch("sources", []).empty?

contents = required.map { |path| root.join(path).read }.join("\n")
template_marker = ["__", "PRODUCT_"].join
abort "unresolved template token" if contents.include?(template_marker)
private_key_pattern = Regexp.new(["-----BEGIN", ".*PRIVATE KEY-----"].join(" "))
abort "possible private key" if contents.match?(private_key_pattern)

puts "Product foundation validation passed"
