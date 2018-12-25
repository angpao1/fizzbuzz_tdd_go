package fizzbuzz_tdd_go

import "strconv"

func FizzBuzz(input int) string {
	switch {
	case input % 15 == 0 : return "fizzbuzz"
	case input % 5 == 0 : return "buzz"
	case input % 3 == 0 : return "fizz"
	default : return strconv.Itoa(input)
	}
}