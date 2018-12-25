package fizzbuzz_tdd_go

import "testing"

func TestInput_1_ShouldBe_1(t *testing.T) {
	fb := FizzBuzz(1)
	expected := "1"
	if fb != expected {
		t.Errorf("Actual = %s, expected = %s", fb, expected)
	}
}