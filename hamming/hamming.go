package hamming

import "errors"

func Distance(a, b string) (int, error) {

	var count int
	if len(a) != len(b) {
		return 0, errors.New("strings must be same length")
	}

	for index := range a {
		if a[index] != b[index] {
			count++
		}
	}
	return count, nil
}
