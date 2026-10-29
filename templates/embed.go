// Package templates embeds the text/template sources for every generated file.
package templates

import "embed"

//go:embed orchestrator sections shared targets runtime
var FS embed.FS
