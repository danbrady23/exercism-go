package collatzconjecture

import "errors"

func CollatzConjecture(n int) (int, error) {
	var count int

	if n < 1 {
		return 0, errors.New("n must be positive int greater than 1")
	}

	for n > 1 {
		if n%2 == 0 {
			n /= 2
		} else {
			n = (3 * n) + 1
		}
		count++
	}
	return count, nil
}
