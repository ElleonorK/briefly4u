// Package schemas compiles the embedded JSON Schemas (draft 2020-12) and validates documents.
package schemas

import (
	"bytes"
	"fmt"
	"sync"

	"github.com/santhosh-tekuri/jsonschema/v6"

	"github.com/ElleonorK/briefly4u/schema"
)

// $id prefix in every schema file: a URN, so the ids don't depend on where the site is hosted
// (<user>.github.io/briefly4u or the maintainer's own domain). Resolution is offline from the embedded FS.
const baseURI = "urn:briefly4u:schema:"

var (
	once     sync.Once
	compiled map[string]*jsonschema.Schema
	initErr  error
)

// Names of the schemas callers can validate against.
var Names = []string{"profile", "pack-topic", "pack-tone", "pack-theme", "pack-faith"}

func load() {
	c := jsonschema.NewCompiler()
	c.DefaultDraft(jsonschema.Draft2020)
	compiled = map[string]*jsonschema.Schema{}
	for _, n := range Names {
		raw, err := schema.FS.ReadFile(n + ".schema.json")
		if err != nil {
			initErr = err
			return
		}
		doc, err := jsonschema.UnmarshalJSON(bytes.NewReader(raw))
		if err != nil {
			initErr = fmt.Errorf("%s: %w", n, err)
			return
		}
		if err := c.AddResource(baseURI+n, doc); err != nil {
			initErr = err
			return
		}
	}
	for _, n := range Names {
		s, err := c.Compile(baseURI + n)
		if err != nil {
			initErr = fmt.Errorf("compile %s: %w", n, err)
			return
		}
		compiled[n] = s
	}
}

// Validate checks v (a value decoded from YAML/JSON into map[string]any etc.) against schema name.
func Validate(name string, v any) error {
	once.Do(load)
	if initErr != nil {
		return initErr
	}
	s, ok := compiled[name]
	if !ok {
		return fmt.Errorf("unknown schema %q", name)
	}
	return s.Validate(v)
}
