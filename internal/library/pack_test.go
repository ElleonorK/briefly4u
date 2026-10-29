package library

import (
	"testing"

	"github.com/ElleonorK/briefly4u/internal/schemas"
)

// Every embedded pack's front matter must validate against pack-<kind>.schema.json.
func TestEmbeddedPacksValidate(t *testing.T) {
	packs, err := Embedded().Packs()
	if err != nil {
		t.Fatal(err)
	}
	if len(packs) < 4 {
		t.Fatalf("want at least one pack per kind, got %d", len(packs))
	}
	seen := map[string]bool{}
	for _, p := range packs {
		seen[p.Kind] = true
		if err := schemas.Validate("pack-"+p.Kind, jsonish(p.FrontMatter)); err != nil {
			t.Errorf("%s: %v", p.Path, err)
		}
	}
	for _, k := range []string{"topic", "tone", "theme", "faith"} {
		if !seen[k] {
			t.Errorf("no %s pack embedded", k)
		}
	}
}

func TestBadPackRejected(t *testing.T) {
	src := []byte("---\napiVersion: briefly4u/v1\nkind: TopicPack\nid: x\nversion: 1.0.0\nlabel: X\ncolour: red\nreaction: r\n---\n# X\n")
	p, err := ParsePack("topics/x.md", src)
	if err != nil {
		t.Fatal(err)
	}
	if schemas.Validate("pack-topic", jsonish(p.FrontMatter)) == nil {
		t.Fatal("colour 'red' accepted; want #rrggbb only")
	}
}
