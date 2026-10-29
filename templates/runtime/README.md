Runtime scripts copied verbatim into every bundle (`scripts/`): assemble.py, render_pdf.sh, xkcd_chip.py.
They run inside the user's agent (Python 3 stdlib + headless Chrome), not in the Go binary.
TODO: port the generic scripts from the reference brief (layout code only, no private details; LOW-LEVEL.md §10) and make the section table data-driven.
