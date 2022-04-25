package goloc

import (
	"testing"

	"golang.org/x/exp/slices"
)

func TestGitignore(t *testing.T) {
	expected := []string{"test", "test2", "test3"}
	gi := LoadGitIgnore("./test_.gitignore")

	if len(gi) != 3 {
		t.Errorf("Expected LoadGitIgnore length to be 3, but got %d", len(gi))
	}

	for _, g := range gi {
		if !slices.Contains(expected, g) {
			t.Errorf("LoadGitIgnore() does not contain the expected values.")
		}
	}
}

func TestLoad(t *testing.T) {
	ignore := []string{"test", "test2", "test3"}
	m := Load("./test_dir", ignore, false)

	if len(m) == 0 {
		t.Errorf("Load() did not load the expected filepaths.")
	}

	if _, ok := m["css"]; !ok {
		t.Errorf("Expected css to exist, but it doesn't.")
	} else {
		if m["css"] != 12 {
			t.Errorf("Expected css LoC to be 12, but got %d", m["css"])
		}
	}

	if _, ok := m["js"]; !ok {
		t.Errorf("Expected js to exist, but it doesn't.")
	} else {
		if m["js"] != 28 {
			t.Errorf("Expected js LoC to be 28, but got %d", m["js"])
		}
	}

	if _, ok := m["html"]; !ok {
		t.Errorf("Expected html to exist, but it doesn't.")
	} else {
		if m["html"] != 12 {
			t.Errorf("Expected html LoC to be 12, but got %d", m["html"])
		}
	}

}
