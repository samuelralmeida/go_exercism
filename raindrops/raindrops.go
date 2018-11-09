package raindrops

import (
	"strconv"
)

// Convert return Pling, Plang, Plong according factors
func Convert(num int) string {
	var resp string

	if num%3 == 0 {
		resp += "Pling"
	}
	if num%5 == 0 {
		resp += "Plang"
	}
	if num%7 == 0 {
		resp += "Plong"
	}

	if resp == "" {
		resp = strconv.Itoa(num)
	}
	return resp
}
