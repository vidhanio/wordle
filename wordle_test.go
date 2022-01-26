package wordle

import (
	"testing"
)

func TestNew(t *testing.T) {
	dictionary := []string{"among", "task", "sus", "lol", "trol"}
	common := []string{"lol", "trol"}

	w, err := New(4, 10, dictionary, common)
	if err != nil {
		t.Errorf("Error: %v", err)
	}

	if len(w.dictionary) != 2 {
		t.Errorf("Expected 2 words in dictionary, got %d", len(w.dictionary))
	}
}
