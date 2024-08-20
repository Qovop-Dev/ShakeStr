package shakestr

import (
	"fmt"
	"math/rand/v2"
	"strings"
)

const (
	WordParam = "w"
	TextParam = "t"
)

// input s = word or text to shake
// input p = parameter (w for word, t for text)
func Shake(s string, p string) (string, error) {

	switch strings.ToLower(p) {
	// word to be shake
	case WordParam:
		return shakeWord(s)

	// text to be shake
	/*case TextParam:

	 */
	// unknown parameter
	default:
		return "", fmt.Errorf("unknown parameter: %s", p)
	}

}

func shakeWord(s string) (string, error) {
	// convert to rune
	runes := []rune(strings.TrimSpace(s))

	// check word lenght
	if len(runes) <= 3 { // Si le mot est trop court, pas besoin de mélanger
		return "", fmt.Errorf("the word length must be greater than 3")
	}

	// Check if it contains any spaces
	if strings.Contains(string(runes), " ") {
		return "", fmt.Errorf("input must be a single word without spaces, got: %s", s)
	}

	// Shake runes between first and last letter
	middle := runes[1 : len(runes)-1]

	//check middle has different letters
	valid := false
	for i := 0; i < len(middle)-1; i++ {
		if middle[i] != middle[i+1] {
			valid = true
			break
		}
	}
	if !valid {
		return "", fmt.Errorf("the word cannot be shuffled because it contains identical adjacent letters in the middle")
	}

	// Copy middle ton check the rand.shuffle result
	original := make([]rune, len(middle))
	copy(original, middle)

	// Loop until the shuffle changes the order of the middle letters
	for {
		rand.Shuffle(len(middle), func(i, j int) {
			middle[i], middle[j] = middle[j], middle[i]
		})

		// Check if middle is different from original
		changed := false
		for i := range middle {
			if middle[i] != original[i] {
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
	shaked := string(runes[0]) + string(middle) + string(runes[len(runes)-1])

	return shaked, nil
}
