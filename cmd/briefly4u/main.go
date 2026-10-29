// Command briefly4u is the Briefly4U CLI: it validates a daily-brief profile and
// generates an agent bundle for a target (grokbot, cursor, claude, generic, actions).
package main

import (
	"os"

	"github.com/ElleonorK/briefly4u/internal/cli"
)

// Set by GoReleaser via -ldflags "-X main.version=... -X main.commit=... -X main.date=...".
var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

func main() {
	if err := cli.NewRoot(cli.BuildInfo{Version: version, Commit: commit, Date: date}).Execute(); err != nil {
		os.Exit(1)
	}
}
