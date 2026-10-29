---
name: briefly4u
description: Briefly4U commands for your daily brief. Use for "/briefly4u setup", "/briefly4u add-section", "/briefly4u tune", "/briefly4u preview", "/briefly4u doctor" and "/briefly4u update", or "set up my daily brief".
license: MIT
---
# Briefly4U

Read the first word after `/briefly4u` and follow the matching file in this plugin exactly:

| Command | Follow |
|---|---|
| `setup` (or no word, or "set up my daily brief") | `../../skills/setup/SKILL.md` |
| `add-section` | `../../skills/add-section/SKILL.md` |
| `tune` | `../../skills/tune/SKILL.md` |
| `preview` | `../../skills/preview/SKILL.md` |
| `doctor` | `../../skills/doctor/SKILL.md` |
| `update` | `../../skills/update/SKILL.md` |

Anything else: list these six commands with one line each.

Why a router: Claude Code namespaces plugin skills (`/briefly4u:setup`), so it reads `skills/` directly. Grok Bot and Cursor install this one skill, so every command starts with `/briefly4u` and no generic name like `/setup` is taken.
