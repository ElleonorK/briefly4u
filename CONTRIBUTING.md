# Contributing to Briefly4U

Thanks! Most contributions are **library packs**: plain files that need no Go.

## Packs
1. Copy an existing pack of the same kind (`library/topics/vuln.md`, `library/tones/cheeky.md`, `library/themes/neon-dark.css`, `library/faiths/jewish.yaml`). Front matter format: LOW-LEVEL.md §4.
2. Rules: **every sample is invented and labelled EXAMPLE**. No real news items (even approved ones), no real people's messages, no private data and no credentials. Sources a pack names must be public feeds or APIs that you checked.
3. Version: patch = typo/sample, minor = new sources/angles/samples, **major = the bar changed**. Add a line to `CHANGELOG.md`.
4. Run `briefly4u lint-packs library/` (schema, required sections, sample labels, tag contrast ≥ 4.5:1 against every theme, no secrets).
5. Open a PR. For now the maintainer reviews every change herself. Area reviewers (CODEOWNERS) across the library come later, once there are regular contributors.
6. New faith packs (for example Hindu, Buddhist or Sikh calendars, not shipped at launch) are welcome once there is a verified, machine-readable source and a reviewer from that tradition.

## Code
- Go 1.23+. `go test ./...`; after an intended output change, `go test ./internal/generator -update` and review the golden diff.
- `golangci-lint run`. Conventional commit prefixes (`feat:`, `fix:`, `pack:`, `docs:`) feed the release notes.
- Changes to `templates/shared/` (safety ground rules) get extra scrutiny in review.

## Support
Community support, no SLA. Bug reports need `/briefly4u doctor` (Claude Code: `/briefly4u:doctor`) or `briefly4u doctor` output with private data removed.

By contributing you agree that your contribution is licensed under the [MIT License](LICENSE), and you follow the [Code of Conduct](CODE_OF_CONDUCT.md).
