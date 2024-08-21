package shakestr

import (
	"fmt"
	"strings"
	"time"

	"golang.org/x/exp/rand"
)

const (
	wordParam    = "w"
	textParam    = "t"
	middleParam  = "m"
	fullParam    = "f"
	reverseParam = "r"
)

// input s = the word or text to shuffle
// input p = parameter ( -> 1st letter = input type : w(word), t(text)
//
//	-> after 1st letter = shake methode :
//	          m(shuffle middle of word),
//	          f(shuffle the full word),
//	          r(reverse the entire word))
func Shake(s string, p string) (string, error) {

	// Check the input is not empty
	if s == "" {
		return "", fmt.Errorf("input value not valid, cannot be empty")
	}

	// the parameter value should be at least 2 caracters
	if len(p) < 2 {
		return "", fmt.Errorf("parameter length must be at least 2, got: %s", p)
	}

	// Assign the parameters
	inputType, shakeMethod := strings.ToLower(string(p[0])), strings.ToLower(string(p[1]))

	switch inputType {
	// Action on single word
	case wordParam:
		switch shakeMethod {
		case middleParam, fullParam:
			return shakeWord(s, shakeMethod)
		case reverseParam:
			return reverse(s)
		default:
			return "", fmt.Errorf("unknown parameter: %s", p)
		}

	// Action on text
	/*case textParam:

	 */
	// unknown parameter
	default:
		return "", fmt.Errorf("unknown parameter: %s", p)
	}

}

func shakeWord(s string, p string) (string, error) {
	// convert to rune
	runes := []rune(strings.TrimSpace(s))

	// buffer variable to work on
	bufferRunes := make([]rune, len(runes))
	copy(bufferRunes, runes)

	// Check if it contains any spaces
	if strings.Contains(string(runes), " ") {
		return "", fmt.Errorf("input must be a single word without spaces, got: %s", s)
	}

	// check word length
	if p == middleParam {
		if len(runes) <= 3 {
			return string(runes), nil
		}
		// Shake runes between first and last letter
		bufferRunes = runes[1 : len(runes)-1]
	} else {
		if len(runes) == 1 {
			return string(runes), nil
		}
	}

	//check bufferRunes has different letters
	valid := false
	for i := 0; i < len(bufferRunes)-1; i++ {
		if bufferRunes[i] != bufferRunes[i+1] {
			valid = true
			break
		}
	}
	if !valid {
		if p == middleParam {
			return "", fmt.Errorf("the word cannot be shuffled because it contains identical adjacent letters in the middle")
		} else {
			return "", fmt.Errorf("the word cannot be shuffled because it contains identical letters")
		}
	}

	// Copy middle to check the rand.shuffle result
	original := make([]rune, len(bufferRunes))
	copy(original, bufferRunes)

	// Loop until the shuffle changes the order of the middle letters
	for {
		rand.Seed(uint64(time.Now().UnixNano()))
		rand.Shuffle(len(bufferRunes), func(i, j int) {
			bufferRunes[i], bufferRunes[j] = bufferRunes[j], bufferRunes[i]
		})

		// Check if bufferRunes is different from original
		changed := false
		for i := range bufferRunes {
			if bufferRunes[i] != original[i] {
				changed = true
				break
			}
		}

		// Break the outer loop if the middle has changed
		if changed {
			break
		}
	}

	// Rebuild the word
	var shaked string
	if p == middleParam {
		shaked = string(runes[0]) + string(bufferRunes) + string(runes[len(runes)-1])
	} else {
		shaked = string(bufferRunes)
	}

	return shaked, nil
}

func reverse(s string) (string, error) {
	// convert to rune
	runes := []rune(strings.TrimSpace(s))
	reverseRunes := make([]rune, len(runes))

	// check word length
	if len(runes) == 1 {
		return string(runes), nil
	}

	// reverse the word
	for i := range runes {
		reverseRunes[len(runes)-i-1] = runes[i]
	}

	return string(reverseRunes), nil
}
