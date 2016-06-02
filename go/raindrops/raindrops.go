package raindrops

import "fmt"

const testVersion = 2

const pling = "Pling"
const plang = "Plang"
const plong = "Plong"

func Convert(num int) string {
	var str string

	if num%3 == 0 {
		str += pling
	}

	if num%5 == 0 {
		str += plang
	}

	if num%7 == 0 {
		str += plong
	}

	if len(str) == 0 {
		str += fmt.Sprintf("%d", num)
	}

	return str
}
