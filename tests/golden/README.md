Golden files: the exact bundle the generator must produce for the invented EXAMPLE profile `examples/profiles/dana.example.yaml`, per target: `dana.example-claude/` (first target, Claude Code) and `dana.example-grokbot/`.
Regenerate after an intended change with `go test ./internal/generator -update`, then review the diff in the PR.
CI fails if a regenerated golden differs from what is committed.
