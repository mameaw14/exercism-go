package etl

const testVersion = 1

func convert(a string) string {
	var margin int
	margin = 'A' - 'a'

	var x int
	x = int(a[0])
	a = string(x - margin)
	return a
}

func Transform(old map[int][]string) map[string]int {
	var n map[string]int
	n = make(map[string]int)
	for v, str := range old {
		for _, a := range str {
			a = convert(a)
			n[a] = v
		}
	}
	return n
}
