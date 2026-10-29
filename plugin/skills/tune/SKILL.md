---
name: tune
description: Change your Briefly4U daily brief in plain words ("fewer launches", "add Berlin", "use °F", "make the closing serious"); shows the profile diff and regenerates only the affected skills. Use for "/briefly4u:tune" (Claude Code) or "/briefly4u tune" (Grok Bot, Cursor).
---
# Tune
1. Translate the request into profile changes (basics first, then settings). The news bar moves only through the `quality` setting (strict, the default, or relaxed); relaxed still requires reliable sources and something new. Never drop those two requirements; offer sources or angles instead.
2. Show the diff and the skills that will be regenerated (from each skill's input hash). On yes, write and regenerate only those.
3. If a generated skill was hand-edited (hash mismatch), show the edit and offer to move it into the profile before overwriting.
