// Package cli wires the cobra command tree for the briefly4u binary.
package cli

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"

	"github.com/ElleonorK/briefly4u/internal/generator"
	"github.com/ElleonorK/briefly4u/internal/library"
	"github.com/ElleonorK/briefly4u/internal/profile"
)

type BuildInfo struct{ Version, Commit, Date string }

func NewRoot(bi BuildInfo) *cobra.Command {
	generator.KitVersion = bi.Version
	var libDir string
	root := &cobra.Command{
		Use:           "briefly4u",
		Short:         "Set up your own daily brief and generate it for your AI agent",
		SilenceUsage:  true,
		SilenceErrors: false,
	}
	root.PersistentFlags().StringVar(&libDir, "library", "", "use a local library dir instead of the packs embedded in this binary")

	src := func() (library.Source, error) {
		if libDir == "" {
			return library.Embedded(), nil
		}
		return library.Dir(libDir)
	}

	root.AddCommand(&cobra.Command{
		Use:   "version",
		Short: "Print version information",
		Run: func(cmd *cobra.Command, _ []string) {
			fmt.Fprintf(cmd.OutOrStdout(), "briefly4u %s (%s, %s)\n", bi.Version, bi.Commit, bi.Date)
		},
	})

	root.AddCommand(&cobra.Command{
		Use:   "validate [profile.yaml]",
		Short: "Validate a profile against the schema and the no-secrets rule",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			b, err := os.ReadFile(args[0])
			if err != nil {
				return err
			}
			_, errs := profile.Parse(b)
			for _, e := range errs {
				fmt.Fprintln(cmd.ErrOrStderr(), "✖", e)
			}
			if len(errs) > 0 {
				return fmt.Errorf("%d problem(s)", len(errs))
			}
			fmt.Fprintln(cmd.OutOrStdout(), "✔ valid")
			return nil
		},
	})

	var target, out string
	var only []string
	build := &cobra.Command{
		Use:   "build [profile.yaml]",
		Short: "Generate the agent bundle for a target",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			b, err := os.ReadFile(args[0])
			if err != nil {
				return err
			}
			p, errs := profile.Parse(b)
			if len(errs) > 0 {
				return fmt.Errorf("profile invalid: %v (run `briefly4u validate`)", errs[0])
			}
			s, err := src()
			if err != nil {
				return err
			}
			bundle, err := generator.Build(p, generator.Options{Target: target, Library: s, Only: only})
			if err != nil {
				return err
			}
			for path, f := range bundle.Files {
				dst := filepath.Join(out, path)
				if err := os.MkdirAll(filepath.Dir(dst), 0o755); err != nil {
					return err
				}
				if err := os.WriteFile(dst, f.Content, 0o644); err != nil {
					return err
				}
			}
			// TODO: write briefly4u.lock (pack ids, versions, hashes, InputHash per file) and --zip.
			fmt.Fprintf(cmd.OutOrStdout(), "✔ %d files → %s\n", len(bundle.Files), out)
			return nil
		},
	}
	build.Flags().StringVar(&target, "target", "claude", "claude|grokbot (implemented); cursor|generic|actions planned")
	build.Flags().StringVar(&out, "out", "dist/bundle", "output directory")
	build.Flags().StringSliceVar(&only, "only", nil, "regenerate only these sections")
	root.AddCommand(build)

	for _, c := range []struct{ use, short string }{
		{"init", "Interview (same questions and defaults as /briefly4u setup) and write profile.yaml"},
		{"config", "Get or set profile values (basics, settings, sections)"},
		{"preview", "Render EXAMPLE fixtures with the chosen theme and tone"},
		{"doctor", "Check Python 3, Chrome, pdfinfo, bundle drift and library versions"},
		{"library", "List, show or scaffold library packs"},
		{"update", "Rebuild against newer packs, showing rule diffs, keeping custom entries"},
		{"migrate", "Upgrade a profile between apiVersions"},
		{"lint-packs", "Lint library packs (schema, samples, contrast, optional URL checks)"},
	} {
		c := c
		root.AddCommand(&cobra.Command{Use: c.use, Short: c.short + " (TODO)", RunE: func(*cobra.Command, []string) error {
			return fmt.Errorf("%s: not implemented yet", c.use)
		}})
	}
	return root
}
