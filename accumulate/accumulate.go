package accumulate

// Accumulate return slice after each element was sent to specific function
func Accumulate(given []string, converter func(string) string) []string {
	var resp []string
	for _, g := range given {
		resp = append(resp, converter(g))
	}
	return resp
}
