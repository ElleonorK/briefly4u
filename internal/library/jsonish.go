package library

import "time"

// jsonish converts YAML-decoded values (ints, timestamps) to the types JSON Schema validation expects.
// TODO: share one implementation with internal/profile.
func jsonish(v any) any {
	switch t := v.(type) {
	case map[string]any:
		m := make(map[string]any, len(t))
		for k, x := range t {
			m[k] = jsonish(x)
		}
		return m
	case []any:
		for i := range t {
			t[i] = jsonish(t[i])
		}
		return t
	case int:
		return float64(t)
	case time.Time:
		return t.Format("2006-01-02")
	default:
		return v
	}
}
