package bob

import (
	"strings"
)

// Hey should have a comment documenting it.
func Hey(remark string) string {
	s := strings.TrimSpace(remark)

	if isSilence(s) {
		return "Fine. Be that way!"
	} else if isQuestion(s) && isYell(s) {
		return "Calm down, I know what I'm doing!"
	} else if isQuestion(s) {
		return "Sure."
	} else if isYell(s) {
		return "Whoa, chill out!"
	}
	return "Whatever."
}

func isSilence(s string) bool {
	return len(s) == 0
}

func isQuestion(s string) bool {
	return s[len(s)-1:] == "?"
}

func isYell(s string) bool {
	return s == strings.ToUpper(s) && s != strings.ToLower(s)
}
