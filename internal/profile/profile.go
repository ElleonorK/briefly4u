// Package profile loads and validates a Briefly4U profile (apiVersion briefly4u/v1, kind Profile).
package profile

import (
	"bytes"
	"fmt"
	"regexp"
	"time"

	"gopkg.in/yaml.v3"

	"github.com/ElleonorK/briefly4u/internal/schemas"
)

const (
	APIVersion = "briefly4u/v1"
	Kind       = "Profile"
)

// Profile is the decoded document. Typed access is added field by field as the
// generator needs it (TODO); Raw keeps the full tree for hashing and templates.
type Profile struct {
	APIVersion string `yaml:"apiVersion"`
	Kind       string `yaml:"kind"`
	Identity   struct {
		Name      string `yaml:"name"`
		AddressAs string `yaml:"address_as"`
	} `yaml:"identity"`
	Tone  string         `yaml:"tone"`
	Theme string         `yaml:"theme"`
	Raw   map[string]any `yaml:"-"`
}

// secretShaped catches values that look like credentials. Secrets must be references
// (secret_ref: "env:NAME", "assistant:<connector>", "keychain:<path>"), never values.
var secretShaped = regexp.MustCompile(`(?i)(xox[abposr]-[0-9A-Za-z-]+|ghp_[0-9A-Za-z]{20,}|sk-[0-9A-Za-z]{20,}|AIza[0-9A-Za-z_-]{30,}|[?&](pwd|passcode|password|pass|pw|token|key|secret|sig|signature)=)`)

// Parse decodes YAML, enforces apiVersion/kind, validates against the JSON Schema
// and rejects secret-shaped values. It returns every problem found, not just the first.
func Parse(src []byte) (*Profile, []error) {
	var errs []error
	var raw map[string]any
	if err := yaml.NewDecoder(bytes.NewReader(src)).Decode(&raw); err != nil {
		return nil, []error{fmt.Errorf("yaml: %w", err)}
	}
	var p Profile
	if err := yaml.Unmarshal(src, &p); err != nil {
		return nil, []error{fmt.Errorf("yaml: %w", err)}
	}
	p.Raw = raw
	if p.APIVersion != APIVersion || p.Kind != Kind {
		errs = append(errs, fmt.Errorf("apiVersion/kind must be %s/%s (got %q/%q); run `briefly4u migrate` for older files", APIVersion, Kind, p.APIVersion, p.Kind))
	}
	if err := schemas.Validate("profile", toJSONTypes(raw)); err != nil {
		errs = append(errs, err)
	}
	if loc := secretShaped.FindIndex(src); loc != nil {
		errs = append(errs, fmt.Errorf("profile contains a secret-shaped value near byte %d; use secret_ref instead", loc[0]))
	}
	return &p, errs
}

// toJSONTypes converts yaml.v3 output (map[string]any, int) to the JSON types the validator expects.
func toJSONTypes(v any) any {
	switch t := v.(type) {
	case map[string]any:
		m := make(map[string]any, len(t))
		for k, x := range t {
			m[k] = toJSONTypes(x)
		}
		return m
	case []any:
		for i := range t {
			t[i] = toJSONTypes(t[i])
		}
		return t
	case int:
		return float64(t)
	case int64:
		return float64(t)
	case time.Time: // YAML timestamps such as 2026-10-20 → JSON Schema "date" strings
		if t.Hour() == 0 && t.Minute() == 0 && t.Second() == 0 {
			return t.Format("2006-01-02")
		}
		return t.Format(time.RFC3339)
	default:
		return v
	}
}
