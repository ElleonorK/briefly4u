// Package schema embeds the JSON Schema files, the language-neutral source of truth
// for profiles and library packs. Non-Go tools (editors, the web builder's form
// renderer, pack CI) read the same .json files directly.
package schema

import "embed"

//go:embed *.schema.json
var FS embed.FS
