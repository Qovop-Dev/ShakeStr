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
		{"Go", "w", "", "parameter length is not valid, got: w"},
		{"Go", "wwww", "", "parameter length is not valid, got: wwww"},
		//shuffle middle of the word
		{"Ace", "wm", "Ace", ""},
		{"(Ace)", "wm", "(Ace)", ""},
		{"(Ace", "wm", "(Ace", ""},
		{"Ace)", "wm", "Ace)", ""},
		{"Ace!", "wm", "Ace!", ""},
		{"Ace !", "wm", "Ace !", ""},
		{"{Ace!}", "wm", "{Ace!}", ""},
		{"Go", "wm", "Go", ""},
		{"Go?", "wm", "Go?", ""},
		{"'Go ?'", "wm", "'Go ?'", ""},
		{"G", "wm", "G", ""},
		{"[G", "wm", "[G", ""},
		{"Test", "wm", "Tset", ""},
		{"Test.", "wm", "Tset.", ""},
		{"Test", "Wm", "Tset", ""},
		{" Test ", "Wm", "Tset", ""},
		{" Test, ", "Wm", "Tset,", ""},
		{"Boot", "Wm", "Boot", "the word cannot be shuffled because it contains identical adjacent letters in the middle"},
		{"Test", "Zm", "", "unknown parameter: Zm"},
		{"Test", "wZ", "", "unknown parameter: wz"},
		{"Invalid Test", "wm", "", "input must be a single word without spaces, got: Invalid Test"},
		//shuffle the full word
		{"G", "wf", "G", ""},
		{"(G)", "wf", "(G)", ""},
		{"G!", "wf", "G!", ""},
		{"G !", "wf", "G !", ""},
		{"Go", "wf", "oG", ""},
		{"(Go", "wf", "(oG", ""},
		{"Go)", "wf", "oG)", ""},
		{"Go?", "wf", "oG?", ""},
		{"Go ?", "wf", "oG ?", ""},
		{"(Go ?)", "wf", "(oG ?)", ""},
		{"Go", "Wf", "oG", ""},
		{" Go ", "Wf", "oG", ""},
		{"oo", "Wf", "oo", "the word cannot be shuffled because it contains identical letters"},
		{"oo!", "Wf", "oo!", "the word cannot be shuffled because it contains identical letters"},
		{"Go", "Zf", "", "unknown parameter: Zf"},
		{"Go", "wZ", "", "unknown parameter: wz"},
		{"Invalid Test", "wf", "", "input must be a single word without spaces, got: Invalid Test"},
		//reverse the full word
		{"A", "wr", "A", ""},
		{"Ace", "wr", "ecA", ""},
		{"Hello You!", "wr", "!uoY olleH", ""},
		{"Hello <You>!", "wr", "!<uoY> olleH", ""},
		{"Hello (You)!", "wr", "!(uoY) olleH", ""},
		{"Hello «You»!", "wr", "!«uoY» olleH", ""},
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

func TestShakeforText(t *testing.T) {
	var tests = []struct {
		s      string
		p      string
		shaked string
		errMsg string
	}{
		{"Hello, it’s nice to meet you!", "tzz", "", "unknown parameter, got: tzz"},
		// reverse mode
		{"Hello, it’s nice to meet you!", "trz", "", "unknown parameter, got: trz"},
		{"Hello, it’s nice to meet you!", "tr", "!uoy teem ot ecin s‘ti ,olleH", ""},
		{"Hello, it’s nice to meet you!", "trp", "!you meet to nice it’s ,Hello", ""},
		{"Hello <You>!", "tr", "!<uoY> olleH", ""},
		{"Hello (You)!", "tr", "!(uoY) olleH", ""},
		{"Hello «You»!", "tr", "!«uoY» olleH", ""},
		//shuffle middle word mode
		{"Give me your gift.", "tm", "Gvie me yuor gfit.", ""},
		{"Help me!", "tmp", "me! Hlep", ""},
		{"Help !", "tmp", "! Hlep", ""},
		{"Help (me!)", "tmp", "(me!) Hlep", ""},
		//shuffle full word mode
		{"He is so up!", "tf", "eH si os pu!", ""},
		{"To me!", "tfp", "em! oT", ""},
		{"To (me!)", "tfp", "(em!) oT", ""},
		{"To !", "tfp", "! oT", ""},
		//shuffle word place
		{"Hello, you!", "tp", "you! Hello,", ""},
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
