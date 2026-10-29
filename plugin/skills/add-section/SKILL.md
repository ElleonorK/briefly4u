---
name: add-section
description: Add one section (or an "add your own" entry) to your Briefly4U daily brief and regenerate only what it affects. Use for "/briefly4u:add-section" (Claude Code) or "/briefly4u add-section" (Grok Bot, Cursor).
---
# Add a section
1. Load `profile.yaml`. Ask which section (or "add your own": source, topic, faith, city, messaging app, joke source).
2. Ask only that section's interview questions (defaults shown). For "add your own", use the generic form: name, source type (feeds and APIs only: RSS, Atom, ICS, JSON API, or a manual list; no web-page scraping) + https URL, auth as a secret reference only, parsing hints, refresh, tag + colour.
3. Show the profile diff; on yes, write it and run `briefly4u build --only <section>` (plus the orchestrator), then replace only those skills.
TODO: per-section question lists.
