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
		{"", "wm", "", "input value not valid, cannot be empty"},
		{"Go", "w", "", "parameter length must be at least 2, got: w"},
		//shuffle middle of the word
		{"Ace", "wm", "Ace", ""},
		{"Go", "wm", "Go", ""},
		{"G", "wm", "G", ""},
		{"Test", "wm", "Tset", ""},
		{"Test", "Wm", "Tset", ""},
		{" Test ", "Wm", "Tset", ""},
		{"Boot", "Wm", "", "the word cannot be shuffled because it contains identical adjacent letters in the middle"},
		{"Test", "Zm", "", "unknown parameter: Zm"},
		{"Test", "wZ", "", "unknown parameter: wZ"},
		{"Invalid Test", "wm", "", "input must be a single word without spaces, got: Invalid Test"},
		//shuffle the full word
		{"G", "wf", "G", ""},
		{"Go", "wf", "oG", ""},
		{"Go", "Wf", "oG", ""},
		{" Go ", "Wf", "oG", ""},
		{"oo", "Wf", "", "the word cannot be shuffled because it contains identical letters"},
		{"Go", "Zf", "", "unknown parameter: Zf"},
		{"Go", "wZ", "", "unknown parameter: wZ"},
		{"Invalid Test", "wf", "", "input must be a single word without spaces, got: Invalid Test"},
		//reverse the full word
		{"A", "wr", "A", ""},
		{"Ace", "wr", "ecA", ""},
		{"Hello You!", "wr", "!uoY olleH", ""},
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
