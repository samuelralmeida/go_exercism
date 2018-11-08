package acronym

import (
	"strings"
	"unicode"
)

// Abbreviate should have a comment documenting it.
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
	result = strings.Replace(result, " ", "", -1)
	result = strings.ToUpper(result)
	return result
}
