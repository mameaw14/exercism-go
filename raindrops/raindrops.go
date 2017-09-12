package raindrops

import "strconv"

const testVersion = 3

func Convert(n int) string {
	var str string
	if n%3 == 0 {
		str += "Pling"
	}
	if n%5 == 0 {
		str += "Plang"
	}
	if n%7 == 0 {
		str += "Plong"
	}
	if str == "" {
		return strconv.Itoa(n)
	}
	return str
}
