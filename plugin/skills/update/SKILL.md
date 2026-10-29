---
name: update
description: Update your Briefly4U daily brief to the library and templates of the installed plugin/binary version; shows rule diffs per pack and keeps your custom entries. Use for "/briefly4u:update" (Claude Code) or "/briefly4u update" (Grok Bot, Cursor).
---
# Update
1. Compare `briefly4u.lock` with the available library. List packs with newer versions; highlight major bumps (the bar changed).
2. Show rule diffs; on yes, rebuild the affected skills, keep `custom:` entries and overrides, write the new lock.
