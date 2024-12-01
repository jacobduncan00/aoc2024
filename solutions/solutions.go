package solutions

import (
	"fmt"
)

func Solve(day int, input string) (interface{}, interface{}) {
	switch day {
	case 1:
		return Day1(input)
	default:
		panic(fmt.Sprintf("Day %d not implemented", day))
	}
}
