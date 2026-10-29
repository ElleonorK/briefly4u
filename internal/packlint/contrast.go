// Package packlint checks library packs: front matter against the pack schemas,
// required body sections, sample labels (REAL+date or EXAMPLE), source URL reachability
// (network, CI only) and tag colour contrast against every theme.
package packlint

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// MinTagContrast is the WCAG AA target for normal-size text (tags render at 16 px).
const MinTagContrast = 4.5

// Contrast returns the WCAG 2.x contrast ratio between two #rrggbb colours.
func Contrast(fg, bg string) (float64, error) {
	l1, err := luminance(fg)
	if err != nil {
		return 0, err
	}
	l2, err := luminance(bg)
	if err != nil {
		return 0, err
	}
	if l1 < l2 {
		l1, l2 = l2, l1
	}
	return (l1 + 0.05) / (l2 + 0.05), nil
}

func luminance(hex string) (float64, error) {
	h := strings.TrimPrefix(hex, "#")
	if len(h) != 6 {
		return 0, fmt.Errorf("colour %q: want #rrggbb", hex)
	}
	var c [3]float64
	for i := 0; i < 3; i++ {
		v, err := strconv.ParseUint(h[2*i:2*i+2], 16, 8)
		if err != nil {
			return 0, fmt.Errorf("colour %q: %w", hex, err)
		}
		s := float64(v) / 255
		if s <= 0.04045 {
			c[i] = s / 12.92
		} else {
			c[i] = math.Pow((s+0.055)/1.055, 2.4)
		}
	}
	return 0.2126*c[0] + 0.7152*c[1] + 0.0722*c[2], nil
}

// TODO: LintDir(fsys) []Finding — walk topics/tones/themes/faiths, validate front matter,
// check samples, contrast of every --tag-* against each theme's --card background,
// and (with -network) HEAD/GET each source URL with a polite timeout.
