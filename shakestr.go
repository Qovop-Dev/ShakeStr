package shakestr

import (
	"fmt"
	"math/rand/v2"
	"strings"
)

type word string

// input s = word or text to shake
// input p = parameter (w for word, t for text)
func Shake(s string, p string) (string, error) {

	switch p {
	// word to be shake
	case "w":
		shaked, err := word(s).shakeWord()
		if err != nil {
			return "", err
		}
		return shaked, err

	// word to be shake
	/*case "t":

	 */
	// unknown parameter
	default:
		return "", fmt.Errorf("unknown parameter")
	}

}

func (s word) shakeWord() (string, error) {
	// convert to rune
	runes := []rune(s)

	// check word lenght
	if len(runes) <= 3 { // Si le mot est trop court, pas besoin de mélanger
		return string(s), fmt.Errorf("to be shaked, the word lenght must be greater than 3")
	}

	// Check if it contains any spaces
	if strings.Contains(string(runes), " ") {
		return string(s), fmt.Errorf("only one word without space required")
	}

	// Shake runes between first and last letter
	middle := runes[1 : len(runes)-1]

	rand.Shuffle(len(middle), func(i, j int) {
		middle[i], middle[j] = middle[j], middle[i]
	})

	// Rebuild the word
	shaked := string(runes[0]) + string(middle) + string(runes[len(runes)-1])
	return shaked, nil
}
