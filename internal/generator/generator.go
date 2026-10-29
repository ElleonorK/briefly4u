// Package generator turns a profile + library packs + embedded templates into an agent bundle.
// Templates are Go text/template with a small allow-listed FuncMap; decisions live here in Go.
package generator

import (
	"bytes"
	"fmt"
	"sort"
	"strings"
	"text/template"

	"github.com/ElleonorK/briefly4u/internal/library"
	"github.com/ElleonorK/briefly4u/internal/profile"
	tpl "github.com/ElleonorK/briefly4u/templates"
)

// KitVersion is stamped into every generated file (overridden by -ldflags at release).
var KitVersion = "0.0.0-dev"

type Options struct {
	Target  string         // claude (first) | grokbot | cursor | generic | actions; chatgpt later
	Library library.Source // embedded by default
	Only    []string       // regenerate only these output ids (sections); empty = all
}

// File is one generated output plus the hash of everything it was derived from.
type File struct {
	Content   []byte
	InputHash string // sha256 over (kit version, template bytes, profile subtrees, pack hashes)
}

type Bundle struct {
	Target string
	Files  map[string]File // bundle-relative path -> file
}

// output declares one generated file: which template, which profile keys it depends on.
type output struct {
	Path      string
	Template  string
	DependsOn []string // dotted profile paths, e.g. "sections.weather", "tone"
}

// targets maps a target to its outputs. Section outputs are expanded per enabled section (TODO).
var targets = map[string][]output{
	// Claude Code is the first target (DECISIONS.md #1): skills under .claude/skills/<name>/SKILL.md.
	"claude": {
		{Path: ".claude/skills/briefly4u-daily-brief/SKILL.md", Template: "orchestrator/SKILL.md.tmpl", DependsOn: []string{"identity", "locale", "calendar_model", "schedule", "delivery", "sections", "tone", "theme"}},
		{Path: ".claude/skills/briefly4u-weather/SKILL.md", Template: "sections/weather/SKILL.md.tmpl", DependsOn: []string{"sections.weather", "locale.units", "tone"}},
		{Path: "README.md", Template: "targets/claude/README.md.tmpl", DependsOn: []string{"identity", "schedule", "delivery", "connections"}},
	},
	"grokbot": {
		{Path: "workflows/briefly4u-daily-brief/SKILL.md", Template: "orchestrator/SKILL.md.tmpl", DependsOn: []string{"identity", "locale", "calendar_model", "schedule", "delivery", "sections", "tone", "theme"}},
		{Path: "workflows/briefly4u-weather/SKILL.md", Template: "sections/weather/SKILL.md.tmpl", DependsOn: []string{"sections.weather", "locale.units", "tone"}},
		{Path: "README.md", Template: "targets/grokbot/README.md.tmpl", DependsOn: []string{"identity", "schedule", "delivery", "connections"}},
	},
}

// funcs is the complete allow-list available to templates. Add functions deliberately.
var funcs = template.FuncMap{
	"join":    join,
	"lower":   strings.ToLower,
	"upper":   strings.ToUpper,
	"enabled": enabled,
	"section": section,
	"dict":    dict,
}

// dict returns an empty map, so templates can default a missing optional block.
func dict() map[string]any { return map[string]any{} }

// join accepts []any (as decoded from YAML) or []string.
func join(v any, sep string) string {
	switch t := v.(type) {
	case []string:
		return strings.Join(t, sep)
	case []any:
		parts := make([]string, len(t))
		for i, x := range t {
			parts[i] = fmt.Sprint(x)
		}
		return strings.Join(parts, sep)
	default:
		return ""
	}
}

// enabled reports whether a section map is on (missing "enabled" means on).
func enabled(s map[string]any) bool {
	v, ok := s["enabled"].(bool)
	return !ok || v
}

// section returns the profile's section with this id, or nil.
func section(p map[string]any, id string) map[string]any {
	list, _ := p["sections"].([]any)
	for _, x := range list {
		if m, ok := x.(map[string]any); ok && m["id"] == id {
			return m
		}
	}
	return nil
}

// Build renders every output for opts.Target.
func Build(p *profile.Profile, opts Options) (*Bundle, error) {
	outs, ok := targets[opts.Target]
	if !ok {
		return nil, fmt.Errorf("target %q not implemented yet (have: %s)", opts.Target, strings.Join(TargetNames(), ", "))
	}
	b := &Bundle{Target: opts.Target, Files: map[string]File{}}
	data := map[string]any{"Profile": p.Raw, "Kit": map[string]any{"Version": KitVersion, "Target": opts.Target}}
	for _, o := range outs {
		src, err := tpl.FS.ReadFile(o.Template)
		if err != nil {
			return nil, err
		}
		t, err := template.New(o.Template).Funcs(funcs).Option("missingkey=error").Parse(string(src))
		if err != nil {
			return nil, fmt.Errorf("%s: %w", o.Template, err)
		}
		// Shared partials ({{template "ground-rules"}}) are parsed into every template set.
		if t, err = t.ParseFS(tpl.FS, "shared/*.md"); err != nil {
			return nil, fmt.Errorf("shared partials: %w", err)
		}
		var buf bytes.Buffer
		if err := t.Execute(&buf, data); err != nil {
			return nil, fmt.Errorf("%s: %w", o.Template, err)
		}
		b.Files[o.Path] = File{Content: buf.Bytes(), InputHash: InputHash(KitVersion, src, p.Raw, o.DependsOn)}
	}
	return b, nil
}

func TargetNames() []string {
	n := make([]string, 0, len(targets))
	for k := range targets {
		n = append(n, k)
	}
	sort.Strings(n)
	return n
}
