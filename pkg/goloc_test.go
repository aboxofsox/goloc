package goloc

import (
	"fmt"
	"testing"

	"golang.org/x/exp/slices"
)

type Test struct {
	key   string
	value int
}

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

func TestExtIgnore(t *testing.T) {
	extignore := []string{
		"html",
		"css",
	}
	tests := []Test{
		{
			key:   "javascript",
			value: 28,
		},
	}
	m := Load("./test_dir", nil, extignore, false)
	fmt.Println(m)
	for _, test := range tests {
		if m[test.key] != test.value {
			t.Errorf("Expected %s LoC to be %d, but got %d", test.key, test.value, m[test.key])
		}
	}

}

func TestLoad(t *testing.T) {
	ignore := []string{"test", "test2", "test3"}
	m := Load("./test_dir", ignore, nil, false)
	tests := []Test{
		{
			key:   "css",
			value: 12,
		},
		{
			key:   "html",
			value: 12,
		},
		{
			key:   "javascript",
			value: 28,
		},
	}

	for _, test := range tests {
		if m[test.key] != test.value {
			t.Errorf("Expected %s LoC to be %d, but got %d", test.key, test.value, m[test.key])
		}
	}

}

func TestLoadIgnore(t *testing.T) {
	gi := LoadGitIgnore("./test_2.gitignore")
	m := Load("./test_dir", gi, nil, false)
	fmt.Println(gi)

	if len(gi) != 1 {
		t.Errorf("Expected LoadGitIgnore() length to be 3, but got %d", len(gi))
	}

	if len(m) > 0 {
		t.Errorf("Expected every file in test_dir to be ignored, but got %d", len(m))
	}
}
