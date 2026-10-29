package generator

import (
	"flag"
	"os"
	"path/filepath"
	"testing"

	"github.com/ElleonorK/briefly4u/internal/library"
	"github.com/ElleonorK/briefly4u/internal/profile"
)

// Run `go test ./internal/generator -update` to rewrite tests/golden after an intended change,
// then review the diff in the PR like any other code change.
var update = flag.Bool("update", false, "rewrite golden files")

// Claude Code is the first target (DECISIONS.md #1); Grok Bot follows with the Cursor plugin format.
func TestGolden(t *testing.T) {
	for _, target := range []string{"claude", "grokbot"} {
		t.Run(target, func(t *testing.T) { golden(t, "dana.example", target) })
	}
}

func golden(t *testing.T, example, target string) {
	src, err := os.ReadFile("../../examples/profiles/" + example + ".yaml")
	if err != nil {
		t.Fatal(err)
	}
	p, errs := profile.Parse(src)
	if len(errs) > 0 {
		t.Fatalf("profile invalid: %v", errs)
	}
	b, err := Build(p, Options{Target: target, Library: library.Embedded()})
	if err != nil {
		t.Fatal(err)
	}
	dir := filepath.Join("../../tests/golden", example+"-"+target)
	for path, f := range b.Files {
		golden := filepath.Join(dir, path)
		if *update {
			_ = os.MkdirAll(filepath.Dir(golden), 0o755)
			if err := os.WriteFile(golden, f.Content, 0o644); err != nil {
				t.Fatal(err)
			}
			continue
		}
		want, err := os.ReadFile(golden)
		if err != nil {
			t.Errorf("%s: missing golden file (run with -update): %v", path, err)
			continue
		}
		if string(want) != string(f.Content) {
			t.Errorf("%s: output differs from golden file; run with -update and review the diff", path)
		}
	}
}

func TestInputHashStable(t *testing.T) {
	p := map[string]any{"tone": "cheeky", "sections": map[string]any{"weather": map[string]any{"cities": []any{"Lisbon"}}}}
	a := InputHash("1.0.0", []byte("t"), p, []string{"tone", "sections.weather"})
	b := InputHash("1.0.0", []byte("t"), p, []string{"tone", "sections.weather"})
	if a != b {
		t.Fatal("hash not deterministic")
	}
	p["tone"] = "serious"
	if a == InputHash("1.0.0", []byte("t"), p, []string{"tone", "sections.weather"}) {
		t.Fatal("hash ignored a dependency change")
	}
	// A key the output doesn't depend on must not change the hash.
	p["tone"] = "cheeky"
	p["theme"] = "neon-dark"
	if a != InputHash("1.0.0", []byte("t"), p, []string{"tone", "sections.weather"}) {
		t.Fatal("hash changed for a non-dependency")
	}
}
