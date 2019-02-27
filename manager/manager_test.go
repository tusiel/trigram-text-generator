package manager

import (
	"testing"
)

func TestGetMapKeys(t *testing.T) {
	chain := make(map[string]map[string]*int)
	chain["One"] = make(map[string]*int)
	chain["Two"] = make(map[string]*int)
	chain["Three"] = make(map[string]*int)
	chain["Four"] = make(map[string]*int)

	m := New()
	m.Chain = chain

	keys := m.getMapKeys()

	for _, expected := range []string{"One", "Two", "Three", "Four"} {
		pass := false
		for _, key := range keys {
			if key == expected {
				pass = true
			}
		}

		if !pass {
			t.Errorf("Expected '%s' to be in the result, but it was not found", expected)
		}
	}
}
