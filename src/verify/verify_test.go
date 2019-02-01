package verify

import "testing"

func TestZoo(t *testing.T) {
	result := Zoo()
	expected := "a"
	if result != expected {
		t.Errorf("got '%s' want '%s'", result, expected)
	}
}
