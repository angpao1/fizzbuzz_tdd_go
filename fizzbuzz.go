package fizzbuzz_tdd_go

import "strconv"

func FizzBuzz(input int) string {
	if input == 5 {
		return "fizz"
	}
	return strconv.Itoa(input)
}