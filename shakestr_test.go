package shakestr

import (
	"testing"
)

func TestShakeforWord(t *testing.T) {
	var tests = []struct {
		s      string
		p      string
		shaked string
		errMsg string
	}{
		{"Ace", "w", "", "the word length must be greater than 3"},
		{"Go", "w", "", "the word length must be greater than 3"},
		{"Test", "w", "Tset", ""},
		{"Test", "W", "Tset", ""},
		{" Test ", "W", "Tset", ""},
		{"Boot", "W", "", "the word cannot be shuffled because it contains identical adjacent letters in the middle"},
		{"Test", "Z", "", "unknown parameter: Z"},
		{"Invalid Test", "w", "", "input must be a single word without spaces, got: Invalid Test"},
	}

	for _, tt := range tests {
		shaked, err := Shake(tt.s, tt.p)

		if shaked != tt.shaked {
			t.Errorf("Invalid result for input (%v, %v). expected=%v, got=%v", tt.s, tt.p, tt.shaked, shaked)
		}

		if tt.errMsg == "" && err != nil {
			t.Errorf("Expected no error for input (%v, %v) but got: %v", tt.s, tt.p, err)
		}

		if tt.errMsg != "" && (err == nil || err.Error() != tt.errMsg) {
			t.Errorf("Invalid error message for input (%v, %v). expected=%v, got=%v", tt.s, tt.p, tt.errMsg, err)
		}

	}
}
