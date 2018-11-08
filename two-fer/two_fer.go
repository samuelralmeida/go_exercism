package twofer

import "fmt"

// Return a short for two for one
func ShareWith(name string) string {
	msg := "One for %s, one for me."
	if name == "" {
		name = "you"
	}
	msg = fmt.Sprintf(msg, name)
	return msg
}
