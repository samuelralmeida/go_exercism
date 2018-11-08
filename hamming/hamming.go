package hamming

import (
	"fmt"
)

func Distance(a, b string) (int, error) {
	count := 0
	if len(a) != len(b) {
		err := fmt.Errorf("diferent length")
		return -1, err
	}
	for i := range a {
		if a[i] != b[i] {
			count += 1
		}
	}
	return count, nil
}
