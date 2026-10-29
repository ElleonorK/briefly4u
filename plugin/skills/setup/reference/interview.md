# Interview script (shared by `/briefly4u setup` (Claude Code: `/briefly4u:setup`), `briefly4u init` and the web builder wizard)

Say first: "I'll ask about 14 short things; say *defaults* at any point to accept the rest."

| # | Ask | Default |
|---|---|---|
| 0 | Confirm detected language · timezone · region · 12/24 h · units | from the agent / OS |
| 1 | Name and what to call you | agent account name |
| 2 | Timezone · work week · first day of week | detected · Mon–Fri (Sun–Thu if region IL) · by region |
| 3 | Delivery time | 08:00 |
| 4 | What counts as "new": since the last successful brief, or a fixed window | since the last successful brief (fixed "10:00 yesterday" when there is no prior brief) |
| 5 | Integration check (no question unless something is missing) | sections without an integration are off |
| 6 | Sections on/off | all with integrations |
| 7 | Weather cities (geocode, confirm by name + country) | the timezone's city |
| 8 | Faiths for Holidays & observances (follow-ups only for picked ones): Jewish, Christian (Catholic, Orthodox new/old calendar, Anglican, Protestant/RCL, Coptic), Muslim, Bahá'í, Pastafarian, national holidays | national holidays of your region, if the country has a verified source; otherwise ask |
| 9 | News topics, order, tag colours (contrast checked) | starter set by reader level |
| 10 | News quality: strict or relaxed | strict (relaxed still needs reliable sources and something new) |
| 11 | Tone (one sample line each) | warm |
| 12 | Theme | calm-minimal |
| 13 | Closing styles · delivery target and artefacts | from tone · agent chat, or Slack DM-to-self if Slack is connected; PDF + phone HTML |
| 14 | Review summary → yes | — |

Integration map: calendar → Google/Microsoft calendar connector (fallback ICS URL) · email → Gmail/Outlook connector (fallback IMAP app password, less safe) · messaging → Slack/Teams connector · weather, observances, news, warning, coming-up, closing → web search + fetch · PDF → shell + headless Chrome (else HTML only) · state files → persistent files (needed for the "since last successful brief" window; without them the fixed window is used).

"Add your own" sources: feeds and APIs only (RSS, Atom, ICS, JSON API) or a manual list. No web-page scraping at launch.
