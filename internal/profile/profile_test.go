package profile

import (
	"os"
	"testing"
)

func TestExamplesValid(t *testing.T) {
	for _, f := range []string{"../../examples/profiles/minimal.yaml", "../../examples/profiles/dana.example.yaml"} {
		b, err := os.ReadFile(f)
		if err != nil {
			t.Fatal(err)
		}
		if _, errs := Parse(b); len(errs) > 0 {
			t.Errorf("%s: %v", f, errs)
		}
	}
}

func TestRejectsSecretValues(t *testing.T) {
	base := "apiVersion: briefly4u/v1\nkind: Profile\nidentity: {name: X}\nlocale: {timezone: UTC}\nsections: []\n"
	cases := map[string]string{
		"inline token":         base + "x-note: xoxb-123456789012-abcdef\n",
		"password key":         base + "connections:\n  - {id: a, provider: imap, auth: app_password, secret_ref: \"env:A\", password: hunter2}\n",
		"secret_ref not a ref": base + "connections:\n  - {id: a, provider: slack, auth: api_key, secret_ref: \"xoxp-1234\"}\n",
		"credential in URL":    base + "custom:\n  - {id: c, kind: source, section: news, name: n, source: {type: rss, url: \"https://e.invalid/f?token=abc\"}}\n",
	}
	for name, doc := range cases {
		if _, errs := Parse([]byte(doc)); len(errs) == 0 {
			t.Errorf("%s: accepted, want rejection", name)
		}
	}
	if _, errs := Parse([]byte(base)); len(errs) > 0 {
		t.Fatalf("base doc should be valid: %v", errs)
	}
}

// DECISIONS.md #2, #3, #5: window modes, news quality levels, and feeds/APIs-only custom sources.
func TestDecisionSettings(t *testing.T) {
	base := "apiVersion: briefly4u/v1\nkind: Profile\nidentity: {name: X}\nlocale: {timezone: UTC}\n"
	valid := map[string]string{
		"window since last success": base + "calendar_model: {window: {mode: since_last_success, max_lookback: 72h}}\nsections: []\n",
		"window fixed":              base + "calendar_model: {window: {mode: fixed, fixed: {from: \"yesterday 10:00\", to: now}}}\nsections: []\n",
		"news relaxed":              base + "sections:\n  - {id: news, basics: {quality: relaxed}}\n",
		"custom rss":                base + "sections: []\ncustom:\n  - {id: c, kind: source, section: news, name: n, source: {type: rss, url: \"https://e.invalid/f.rss\"}}\n",
	}
	invalid := map[string]string{
		"unknown window mode":  base + "calendar_model: {window: {mode: rolling}}\nsections: []\n",
		"unknown news quality": base + "sections:\n  - {id: news, basics: {quality: loose}}\n",
		"custom html scraping": base + "sections: []\ncustom:\n  - {id: c, kind: source, section: news, name: n, source: {type: html_page, url: \"https://e.invalid/\"}}\n",
	}
	for name, doc := range valid {
		if _, errs := Parse([]byte(doc)); len(errs) > 0 {
			t.Errorf("%s: rejected: %v", name, errs)
		}
	}
	for name, doc := range invalid {
		if _, errs := Parse([]byte(doc)); len(errs) == 0 {
			t.Errorf("%s: accepted, want rejection", name)
		}
	}
}
