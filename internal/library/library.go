// Package library selects the pack source: the packs embedded in the binary, or a local
// directory (--library) such as a newer checkout of the repo's library/ folder.
package library

import (
	"io/fs"
	"os"

	lib "github.com/ElleonorK/briefly4u/library"
)

// Source is any fs.FS laid out like library/ (topics/, tones/, themes/, faiths/).
type Source struct {
	FS     fs.FS
	Origin string // "embedded" or the directory path, recorded in briefly4u.lock
}

func Embedded() Source { return Source{FS: lib.FS, Origin: "embedded"} }

func Dir(path string) (Source, error) {
	if _, err := os.Stat(path); err != nil {
		return Source{}, err
	}
	return Source{FS: os.DirFS(path), Origin: path}, nil
}

// TODO: Pack(kind, id) (*Pack, error) — parse YAML front matter, validate against pack-<kind>.schema.json,
// return front matter + body + sha256 for the lockfile.
