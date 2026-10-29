package library

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io/fs"
	"path"
	"strings"

	"gopkg.in/yaml.v3"
)

// Pack is one library file: its front matter, its body, and a content hash for the lockfile.
type Pack struct {
	Kind        string // topic | tone | theme | faith
	Path        string
	FrontMatter map[string]any
	Body        []byte
	SHA256      string
}

var dirKinds = map[string]string{"topics": "topic", "tones": "tone", "themes": "theme", "faiths": "faith"}

// ParsePack splits front matter from body. Formats:
//
//	.md   ---\n<yaml>\n---\n<markdown>
//	.css  /*---\n<yaml>\n---*/\n<css variables>
//	.yaml the whole file is the front matter (faith packs)
func ParsePack(p string, src []byte) (*Pack, error) {
	kind := dirKinds[strings.Split(p, "/")[0]]
	sum := sha256.Sum256(src)
	pk := &Pack{Kind: kind, Path: p, SHA256: hex.EncodeToString(sum[:])}
	var fm []byte
	switch path.Ext(p) {
	case ".md":
		fm, pk.Body = split(src, "---\n", "\n---\n")
	case ".css":
		fm, pk.Body = split(src, "/*---\n", "\n---*/\n")
	case ".yaml", ".yml":
		fm = src
	default:
		return nil, fmt.Errorf("%s: unsupported pack file type", p)
	}
	if fm == nil {
		return nil, fmt.Errorf("%s: missing front matter", p)
	}
	if err := yaml.Unmarshal(fm, &pk.FrontMatter); err != nil {
		return nil, fmt.Errorf("%s: front matter: %w", p, err)
	}
	return pk, nil
}

func split(src []byte, open, close string) (fm, body []byte) {
	if !bytes.HasPrefix(src, []byte(open)) {
		return nil, src
	}
	rest := src[len(open):]
	i := bytes.Index(rest, []byte(close))
	if i < 0 {
		return nil, src
	}
	return rest[:i], rest[i+len(close):]
}

// Packs lists every pack in a source, skipping README-style files that start with "_".
func (s Source) Packs() ([]*Pack, error) {
	var out []*Pack
	err := fs.WalkDir(s.FS, ".", func(p string, d fs.DirEntry, err error) error {
		if err != nil || d.IsDir() || strings.HasPrefix(path.Base(p), "_") || path.Ext(p) == ".go" {
			return err
		}
		if _, ok := dirKinds[strings.Split(p, "/")[0]]; !ok {
			return nil
		}
		src, err := fs.ReadFile(s.FS, p)
		if err != nil {
			return err
		}
		pk, err := ParsePack(p, src)
		if err != nil {
			return err
		}
		out = append(out, pk)
		return nil
	})
	return out, err
}
