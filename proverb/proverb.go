package proverb

import "fmt"

// Proverb should have a comment documenting it.
func Proverb(rhyme []string) []string {
	resp := []string{}
	length := len(rhyme)
	if length == 0 {
		return resp
	}
	for i := range rhyme {
		if i == length-1 {
			resp = append(resp, fmt.Sprintf("And all for the want of a %s.", string(rhyme[0])))
		} else {
			resp = append(resp, fmt.Sprintf("For want of a %s the %s was lost.", rhyme[i], rhyme[i+1]))
		}
	}
	return resp

}
