---
apiVersion: briefly4u/v1
kind: TopicPack
id: vuln
label: VULN
colour: "#ffc53d"   # amber
version: 1.0.0
added: 2026-10-20
reaction: "I need to patch this today"
fun: false
see_also: [breach]
open_story_log: true
sources:
  trusted: [UK NCSC, CERT-EU, ENISA EUVD, Israel National Cyber Directorate, JPCERT/CC, ACSC (au), BSI (de), CERT-FR, BleepingComputer, The Record, Help Net Security, SecurityWeek, Microsoft MSRC, Apple, Chrome releases, Fortinet PSIRT, Ivanti, Cisco, watchTowr Labs, Rapid7, Horizon3, Project Zero, Wiz, Volexity, Shadowserver]
  leads_only: [CISA KEV (one input only), Hacker News, r/netsec, X news search]
status: stable
license: MIT
maintainers: []   # the maintainer reviews all changes for now (DECISIONS.md #12)
checked: 2026-10-22
---
# VULN

## What qualifies (the bar)
A flaw that is actually being exploited or trivially exploitable in software people really run, so it has to be patched or mitigated now. Judged on real-world exploitation and blast radius, not CVSS score or which government list it's on. For vulns and breaches, say concretely what an attacker can do and who is affected. A national agency officially confirming active exploitation inside the window (e.g. a KEV addition, a CERT alert) counts as a new development for an ONGOING item.

## Not
Theoretical bugs in niche products, vendor PR, a CVSS score or a government list on its own.

## Search angles
1. Non-US agencies: UK NCSC news, CERT-EU advisories, ENISA EUVD, Israel National Cyber Directorate, JPCERT/CC, ACSC (au), BSI (de), CERT-FR
2. BleepingComputer, The Record, Help Net Security, SecurityWeek
3. "actively exploited" / "exploited in the wild" / "emergency patch" / "out-of-band"
4. Researcher blogs: watchTowr Labs, Rapid7, Horizon3, Project Zero, Wiz, Volexity, Shadowserver exposure counts
5. CISA KEV as one input only
6. Vendor advisories: Microsoft MSRC, Apple, Chrome releases, Fortinet PSIRT, Ivanti, Cisco
7. Hacker News + r/netsec top of day; X news on zero-day

## Sample items
**Good — EXAMPLE (invented, ONGOING):**
- Title: *"<Product> server: a no-password bug chain is hijacking unpatched machines"*
- Summary: *who runs it, the fixed version, what the attacker gets, and the new development (vendor confirmed widespread exploitation); end with the action: update, and treat older exposed servers as compromised.*

**Bad — EXAMPLE:** *"CVSS 9.8 bug in <niche plugin>, no exploitation seen"* → score alone, niche.
