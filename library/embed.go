// Package library embeds the default library packs into the binary.
// Users can point the CLI at a newer local checkout with --library <dir>.
package library

import "embed"

//go:embed topics tones themes faiths
var FS embed.FS
