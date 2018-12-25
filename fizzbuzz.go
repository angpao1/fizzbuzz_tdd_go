package fizzbuzz_tdd_go

import "strconv"

func FizzBuzz(input int) string {
	if input % 5 == 0{
		return "buzz"
	}
	if input % 3 == 0{
		return "fizz"
	}
	return strconv.Itoa(input)
}