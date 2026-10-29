# Security policy

Briefly4U is a generator: it writes prompts, skills and scripts that run in **your** AI agent with **your** connectors. The most important class of problem is therefore **prompt injection**: content the brief gathers (an email, a Slack message, a web page, a calendar invite, a feed) making the agent do something other than build the brief.

## Please report privately
Use GitHub's private vulnerability reporting: open this repository's **Security** tab and choose **Report a vulnerability**. Don't open a public issue.

## In scope
- A gathered input that makes a generated skill send, post, reply, delete, change settings, or reveal data beyond the user's own brief.
- Generated output that renders a credential-bearing link, unescaped HTML, scripts or styles from gathered content.
- The generator, plugin or builder accepting a secret value in a profile or pack, or writing one to disk.
- Supply chain: release artifacts, signatures, SBOMs, package-manager manifests, CI workflows.

## Out of scope
- Behaviour of the agent platforms themselves (report to their vendors), unless our templates make it worse.
- News quality, joke quality and wrong holiday dates (open a normal issue or a pack PR).

## What to include
Target and version (`briefly4u version`), the profile with private data removed, the gathered input that triggered it (redacted), and what the agent did.

## Response
Best effort, community project with no SLA: acknowledgement within 7 days, fix or mitigation and an advisory as soon as practical. Fixes to the ground rules in `templates/shared/` ship in a patch release and are flagged in `/briefly4u update`.
