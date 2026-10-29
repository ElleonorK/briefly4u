package generator

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"strings"
)

// InputHash identifies everything an output depends on. If it is unchanged, the output
// is unchanged, so `/briefly4u tune` (Claude Code: `/briefly4u:tune`) and `briefly4u build --only` can skip it.
// encoding/json sorts map keys, which makes the serialisation canonical.
func InputHash(kitVersion string, template []byte, profile map[string]any, dependsOn []string) string {
	h := sha256.New()
	h.Write([]byte("briefly4u-input-hash/v1\x00" + kitVersion + "\x00"))
	h.Write(template)
	for _, path := range dependsOn {
		h.Write([]byte("\x00" + path + "="))
		v, _ := json.Marshal(lookup(profile, path))
		h.Write(v)
	}
	// TODO: include the content hash of every library pack the output inlines (from briefly4u.lock).
	return "sha256:" + hex.EncodeToString(h.Sum(nil))
}

func lookup(m map[string]any, dotted string) any {
	var cur any = m
	for _, k := range strings.Split(dotted, ".") {
		mm, ok := cur.(map[string]any)
		if !ok {
			return nil
		}
		cur = mm[k]
	}
	return cur
}
