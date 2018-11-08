package raindrops

import (
	"strconv"
)

// Convert return Pling, Plang, Plong according factors
func Convert(num int) string {
	// store num to return if necessary
	n := num

	// store factors in map to search fastly
	factors := map[int]bool{}

	// first divide num until turn it in odd
	for num%2 == 0 {
		factors[2] = true
		num = num / 2
	}

	// if num is odd, we can try divid it for odds too
	// so we can skip two by two
	for i := 3; i*i <= num; i = i + 2 {
		for num%i == 0 {
			factors[i] = true
			num = num / i
		}
	}

	// ifnum not divide for odds, so it is prime
	if num > 2 {
		factors[num] = true
	}

	resp := checkFactors(factors)
	if resp == "" {
		resp = strconv.Itoa(n)
	}
	return resp
}

func checkFactors(factors map[int]bool) string {
	var resp string
	if factors[3] {
		resp += "Pling"
	}
	if factors[5] {
		resp += "Plang"
	}
	if factors[7] {
		resp += "Plong"
	}
	return resp
}
