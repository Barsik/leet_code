package p292

import (
	"testing"
)

func TestCanWinNim(t *testing.T) {
	result := canWinNim(2)
	if !result {
		t.Error("Expected true")
	}
	result = canWinNim(8)
	if result {
		t.Error("Expected false")
	}
	result = canWinNim(10)
	if !result {
		t.Error("Expected true")
	}
}
