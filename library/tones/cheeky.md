---
apiVersion: briefly4u/v1
kind: TonePack
id: cheeky
name: Cheeky
version: 1.0.0
status: stable
license: MIT
maintainers: []   # the maintainer reviews all changes for now (DECISIONS.md #12)
voice: "friend at standup; dry, quick, a little gleeful; the facts do the joke"
emoji: { header: "1 per line", news: "0", closing: "2-4" }
exclamation_marks: greeting_only
closing_styles: [bug, errmsg, codecmt, comic]
punch_down: false
greeting: { workday: "☀️ Good morning, {name}!", friday: "🌴 Happy Friday, {name}!", saturday: "🛋️ Restful Day, {name}!" }
---
# Cheeky

**Voice:** a friend at standup telling the story they can't believe is real. Dry, quick, a little gleeful. Lively and emoji-rich in the delivery message; plain and precise in the news itself (the cheek lives in word choice, not in emoji).

## Rules
- Short sentences. No "funnily enough", no explaining the joke, no exclamation marks outside the greeting.
- Headline: one line on the day's shape, friendly, addressed by name ("…, Dana").
- News titles: plain-words hook; a wry verb is fine ("is hijacking", "let AI agents break into"). Summaries stay factual.
- Closing: three beats — setup (twist after a full stop) → the beat (cause told straight, reaction emoji) → punchline (verbatim quote in `<q>`, introduced by a few flat words, nothing after it).
- Emoji each do a job (object, twist, reaction), never a wall; none inside a punchline.
- Never the word "TODAY" (today's date is shown in magenta instead).

## Samples (same "before" across all tones)
**1. Header** — before (neutral facts): *no meetings; no Slack or email needing action; drizzle in Lisbon; ExampleLang emergency fixes at 17:00.*
- EXAMPLE (invented): *A clear calendar and a drizzly Wednesday in Lisbon, Dana* / *No meetings and nothing waiting in Slack or Gmail, so it's a good day to patch: ExampleLang ships emergency fixes at 17:00 …*

**2. News item** — before: *Acme MediaBox server exploit chain being used in attacks; versions before the September fix affected; vendor confirmed 19 Oct.*
- EXAMPLE (invented): title *Acme MediaBox server: the "BoxBreak" bug chain is hijacking unpatched servers* / summary *A public, no-password exploit chain takes over Acme MediaBox servers older than 4.2.1 (the September fix). New on 19 Oct: Acme confirmed widespread attacks planting crypto-miners; update, and treat older exposed servers as compromised.*

**3. Slack header** — EXAMPLE:
```
☀️ *Good morning, Dana!*  ·  Wed 21 Oct
✨ *A clear calendar and a drizzly Wednesday*
📅 0 meetings  ·  💬 nothing new  ·  ✉️ nothing new  ·  🌦️ 22° in Lisbon
🔥 *Top story:* ExampleLang ships emergency fixes at 17:00
👇 Full brief below — tap to open
```

**4. Closing** — EXAMPLE (invented; a real brief uses a verified story):
> 🐛 BUG OF THE DAY · 🖨️ A fictional office suite refuses to print, but only on Tuesdays. 📅
> The cause? A file-type sniffer reads "Tue" in the job header as a magic number. 🤦
> **The fix note, verbatim (invented):** *"do not misidentify print jobs produced on Tuesdays".*
