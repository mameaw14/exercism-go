package hamming

import "errors"

const testVersion = 6

func Distance(a, b string) (int, error) {
	if a == b {
		return 0, nil
	}
	if len(a) != len(b) {
		return -1, errors.New("disallow input")
	}

	value, _ := Distance(a[1:], b[1:])
	eq := 1
	if a[0] == b[0] {
		eq = 0
	}
	return eq + value, nil
	return 1, nil
}
