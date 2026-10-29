---
name: doctor
description: Check your Briefly4U setup (integrations per section, schedule, state files, Python 3 and Chrome for PDF, drift between installed skills and the profile, library versions). Use for "/briefly4u:doctor" (Claude Code) or "/briefly4u doctor" (Grok Bot, Cursor).
---
# Doctor
1. List tools and connectors; report per section: ready / missing (with fallback).
2. Confirm the schedule exists and names the delivery target.
3. With a shell: `briefly4u doctor --target <this agent>` (Python 3, Chrome, pdfinfo, lockfile vs library, input-hash drift).
4. `--send-test` only on explicit request: send one test brief to the user's own target and record `confirmed_at`.
Never read message or mail contents to "test" a connection.
