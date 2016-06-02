package hamming

import "errors"

const testVersion = 4

func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return -1, errors.New("Strings must be equal")
	}

	distance := 0
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			distance++
		}
	}

	return distance, nil
}
