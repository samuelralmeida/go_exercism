package acronym

import (
	"unicode"
)

func Abbreviate(s string) string {
	var result string
	for i, c := range s {
		if i == 0 {
			result = result + string(c)
		} else if i == len(s)-1 {
			// break loop if last char is not a letter
			break
		} else if !unicode.IsLetter(c) {
			result = result + string(s[i+1])
		}
	}
	return result
}
