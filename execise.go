package golang

import (
	"strconv"
)

func getCaptcha(v []int) string {
	mode := v[0]
	num := v[1]
	operation := v[2]
	str := v[3]

	s := ""

	data := []string{"one", "two"}
	ope := []string{"+", "-", "*"}

	if mode == 1 {
		s += strconv.Itoa(num) + " " + ope[operation-1] + " " + data[str-1]
	} else {
		s += data[str-1] + " " + ope[operation-1] + " " + strconv.Itoa(num)
	}

	return s
}
