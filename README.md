# Briefly4U

**Your own daily brief, run by your own AI agent.** Free and open source (MIT). No account, nothing hosted, no credentials leave your agent.

> Status: pre-release scaffold (v0.1 in progress). Everything labelled EXAMPLE is invented sample data.

<!-- TODO: screenshot of an invented EXAMPLE brief (neon-dark, laptop + phone) -->

## What you get
One morning page: calendar, messages and mail that need you, weather, holidays and observances, a short high-bar news list, updates that are breaking devices, upcoming events, and a closing joke, in the tone and theme you choose. Delivered as a PDF and a phone-friendly HTML page wherever your agent can reach you.

## Get started (pick one)
1. **Plugin (recommended):** install the Briefly4U plugin, then run setup:
   - Claude Code (first target): `/plugin marketplace add <user>/briefly4u`, `/plugin install briefly4u@briefly4u`, then `/briefly4u:setup`.
   - Grok Bot and Cursor: install the same plugin, then type `/briefly4u setup`.
2. **Web builder:** `https://<user>.github.io/briefly4u/builder/` (or the maintainer's own domain): build `profile.yaml` and a bundle zip in your browser.
3. **CLI** (package names are registered at the first release; nothing is published yet):
   ```sh
   brew install <user>/tap/briefly4u           # macOS / Linux (Homebrew formula; the only macOS offering)
   scoop bucket add briefly4u https://github.com/<user>/scoop-bucket && scoop install briefly4u
   winget install <id-chosen-at-first-release>  # Windows
   yay -S briefly4u-bin                         # Arch (AUR)
   # .deb / .rpm / .apk and archives: GitHub Releases
   briefly4u init && briefly4u build --target claude
   ```

## How it works
`profile.yaml` + library packs → `briefly4u` generator → an agent bundle (orchestrator + one skill per section + scripts + theme) → your agent runs it on its own schedule. Design: [DESIGN.md](docs/DESIGN.md) · implementation notes: [LOW-LEVEL.md](docs/LOW-LEVEL.md). TODO: copy the design docs into docs/.

## Supported agents
| Target | Status |
|---|---|
| Claude Code | first target (v0.1) |
| Grok Bot, Cursor | v0.2 (same plugin, Cursor plugin format) |
| Generic markdown bundle, GitHub Actions | planned (v0.3) |
| ChatGPT / OpenAI | later, as a proper plugin (format verified when scheduled) |

## Safety
Gathered content is treated as data, never instructions; everything is HTML-escaped; credential-bearing links are never rendered; section skills are read-only; the brief goes only to your own confirmed targets. Report problems privately: [SECURITY.md](SECURITY.md).

## Verify a release
```sh
cosign verify-blob --bundle checksums.txt.sigstore.json \
  --certificate-identity "https://github.com/<user>/briefly4u/.github/workflows/release.yml@refs/tags/vX.Y.Z" \
  --certificate-oidc-issuer https://token.actions.githubusercontent.com checksums.txt
sha256sum --check --ignore-missing checksums.txt
```

## Contributing
Packs (topics, tones, themes, faiths) are plain files and welcome as PRs. See [CONTRIBUTING.md](CONTRIBUTING.md). Community support, no SLA; the maintainer reviews every change for now.

## License
[MIT](LICENSE) for everything: code, templates and library pack text. Copyright (c) 2026 Elleonor Kulesko.
