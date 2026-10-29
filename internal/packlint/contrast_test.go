package packlint

import "testing"

func TestContrast(t *testing.T) {
	r, err := Contrast("#000000", "#ffffff")
	if err != nil || r < 20.9 || r > 21.1 {
		t.Fatalf("black/white = %v, %v; want 21", r, err)
	}
	// light-print tag-android, adjusted in the design phase to pass 4.5:1 on white
	if r, _ := Contrast("#18793f", "#ffffff"); r < MinTagContrast {
		t.Fatalf("#18793f on white = %.2f; want >= %.1f", r, MinTagContrast)
	}
}
