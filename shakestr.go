package shakestr

import (
	"fmt"
	"strings"
	"sync"
	"time"
	"unicode"

	"golang.org/x/exp/rand"
)

const (
	wordParam    = "w"
	textParam    = "t"
	middleParam  = "m"
	fullParam    = "f"
	reverseParam = "r"
	placeParam   = "p"
)

//-----------------------------------------------------------------------//
//-----------------------------------------------------------------------//
// input s = the word or text to shuffle
// input p = parameter ( -> 1st letter = input type : w(word), t(text)
//
//	-> after 1st letter for word type input = shake method :
//	          m(shuffle middle of word),
//	          f(shuffle the full word),
//	          r(reverse the entire word))

//	-> after 1st letter for text type input = shake method :
//	          r(reverse the whole text),
//	          rp(reverse only the place of words in the text),
//	          f(shuffle every word (entire) but keep the order)
//	          fp(shuffle every word (entire) and the order),
//	          m(shuffle every word (middle) but keep the order),
//	          mp(shuffle every word (middle) and the order),
//	          p(shuffle only the order))
//-----------------------------------------------------------------------//
//-----------------------------------------------------------------------//

func Shake(s, p string) (string, error) {

	//Validate Inputs
	if err := validateInputs(s, p); err != nil {
		return "", err
	}

	// Assign the parameters
	inputType, shakeMethod, shakeMethodOpt := parseParameters(p)

	switch inputType {
	case wordParam:
		// Action on single word
		return wordProcess(s, shakeMethod)

	case textParam:
		// Action on text
		return textProcess(s, shakeMethod, shakeMethodOpt)

	default:
		// unknown parameter
		return "", fmt.Errorf("unknown parameter: %s", p)
	}

}

//-----------------------------------------//
//---------  Principal Functions  ---------//
//-----------------------------------------//

func validateInputs(s, p string) error {
	// Check the input is not empty
	if s == "" {
		return fmt.Errorf("input value not valid, cannot be empty")
	}

	// the parameter value should be between 2 and 3 caracters
	if len(p) < 2 || len(p) > 3 {
		return fmt.Errorf("parameter length is not valid, got: %s", p)
	}

	return nil
}

func parseParameters(p string) (inputType, shakeMethod, shakeMethodOpt string) {
	inputType = strings.ToLower(string(p[0]))
	shakeMethod = strings.ToLower(string(p[1]))
	if len(p) == 3 {
		shakeMethodOpt = strings.ToLower(string(p[2]))
	}
	return
}

func wordProcess(s, shakeMethod string) (string, error) {
	switch shakeMethod {
	case middleParam, fullParam:
		return shakeWord(s, shakeMethod)
	case reverseParam:
		return reverse(s)
	default:
		return "", fmt.Errorf("unknown parameter: w%s", shakeMethod)
	}
}

func textProcess(s, shakeMethod, shakeMethodOpt string) (string, error) {
	if shakeMethod == reverseParam && shakeMethodOpt == "" {
		return reverse(s)
	} else {
		return shakeText(s, shakeMethod, shakeMethodOpt)
	}
}

//-----------------------------------------//
//--------- Operational Functions ---------//
//-----------------------------------------//

func shakeText(s string, shakeMethod string, shakeMethodOpt string) (string, error) {
	// cut text in slice of words
	words := []string(strings.Fields(strings.TrimSpace(s)))

	// check words slice length
	if len(words) == 1 {
		return "", fmt.Errorf("input must be more than one word, got: %s", strings.Join(words, ""))
	}

	switch shakeMethod {

	case middleParam, fullParam:
		shakedWords := make([]string, len(words))
		// init waitgroup
		var wg sync.WaitGroup

		// launch every word in go routine
		for i := range words {
			wg.Add(1)
			go func(i int) {
				defer wg.Done()
				shakedWord, _ := shakeWord(words[i], shakeMethod)
				shakedWords[i] = shakedWord
			}(i)
		}
		// wait the end of process
		wg.Wait()

		//shake place of word is asked
		if shakeMethodOpt == placeParam {
			return shakePlace(shakedWords)
		} else {
			return strings.Join((shakedWords), " "), nil
		}

	case placeParam:
		return shakePlace(words)

	case reverseParam:
		if shakeMethodOpt == placeParam {
			return reversePlace(words)
		} else {
			return "", fmt.Errorf("unknown parameter, got: %s", ("t" + shakeMethod + shakeMethodOpt))
		}

	default:
		return "", fmt.Errorf("unknown parameter, got: %s", ("t" + shakeMethod + shakeMethodOpt))
	}
}

func shakeWord(s string, p string) (string, error) {
	// convert to rune
	runes := []rune(strings.TrimSpace(s))

	// buffer variable to work on
	var bufferRunes []rune

	//exclude ponctuation
	ponctuation := ""
	ponc := 0
	if unicode.IsPunct(runes[len(runes)-1]) {
		// buffer variable to work on
		bufferRunes = make([]rune, len(runes)-1)
		// separating the ponctuation from the word
		ponc = 1
		ponctuation = string(runes[len(runes)-1])
		copy(bufferRunes, runes[:(len(runes)-1)])
	} else {
		// buffer variable to work on
		bufferRunes = make([]rune, len(runes))
		copy(bufferRunes, runes)
	}

	// Check if it contains any spaces
	if strings.Contains(string(runes), " ") {
		return "", fmt.Errorf("input must be a single word without spaces, got: %s", s)
	}

	// check word length
	if p == middleParam {
		if len(runes) <= (3 + ponc) {
			return string(runes), nil
		}
		// Shake runes between first and last letter
		bufferRunes = runes[1 : len(runes)-1-ponc]
	} else {
		if (len(runes) == 1) || (len(runes) == 2 && ponc == 1) {
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
			return string(runes), fmt.Errorf("the word cannot be shuffled because it contains identical adjacent letters in the middle")
		} else {
			return string(runes), fmt.Errorf("the word cannot be shuffled because it contains identical letters")
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
		shaked = string(runes[0]) + string(bufferRunes) + string(runes[len(runes)-1-ponc]) + ponctuation
	} else {
		shaked = string(bufferRunes) + ponctuation
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

func reversePlace(w []string) (string, error) {
	// result variable
	reverseWords := make([]string, len(w))

	// reverse the word place
	for i := range w {
		//if last character is a ponctuation, then move the ponctuation to the other side of the word
		if unicode.IsPunct(rune(w[i][len(w[i])-1])) {
			reverseWords[len(w)-i-1] = string(w[i][len(w[i])-1]) + string(w[i][:len(w[i])-1])
		} else {
			reverseWords[len(w)-i-1] = w[i]
		}
	}

	return strings.Join((reverseWords), " "), nil
}

func shakePlace(w []string) (string, error) {

	// Copy input to check later the rand.shuffle result
	shakedPlace := make([]string, len(w))
	copy(shakedPlace, w)

	// Loop until the shuffle changes the order of the words place
	for {
		rand.Seed(uint64(time.Now().UnixNano()))
		rand.Shuffle(len(shakedPlace), func(i, j int) {
			shakedPlace[i], shakedPlace[j] = shakedPlace[j], shakedPlace[i]
		})

		// Check if bufferRunes is different from original
		changed := false
		for i := range shakedPlace {
			if shakedPlace[i] != w[i] {
				changed = true
				break
			}
		}

		// Break the outer loop if the words place has changed
		if changed {
			break
		}
	}

	return strings.Join((shakedPlace), " "), nil
}
