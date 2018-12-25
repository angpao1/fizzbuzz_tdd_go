package fizzbuzz_tdd_go

import "testing"

func TestInput_1_ShouldBe_1(t *testing.T) {
	fb := FizzBuzz(1)
	expected := "1"
	if fb != expected {
		t.Errorf("Actual = %s, expected = %s", fb, expected)
	}
}

func TestInput_2_ShouldBe_2(t *testing.T) {
	fb := FizzBuzz(2)
	expected := "2"
	if fb != expected {
		t.Errorf("Actual = %s, expected = %s", fb, expected)
	}
}
func TestInput_3_ShouldBe_3(t *testing.T) {
	fb := FizzBuzz(3)
	expected := "3"
	if fb != expected {
		t.Errorf("Actual = %s, expected = %s", fb, expected)
	}
}