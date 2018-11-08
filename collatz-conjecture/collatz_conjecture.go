package collatzconjecture

import (
	"fmt"
)

func CollatzConjecture(n int) (int, error) {
	if n <= 0 {
		err := fmt.Errorf("n must be greater than zero")
		return 0, err
	} else if n == 1 {
		return 0, nil
	} else if n%2 == 0 {
		step, err := CollatzConjecture(n / 2)
		if err != nil {
			return 0, err
		}
		return 1 + step, nil
	} else {
		step, err := CollatzConjecture(3*n + 1)
		if err != nil {
			return 0, err
		}
		return 1 + step, nil
	}
}
