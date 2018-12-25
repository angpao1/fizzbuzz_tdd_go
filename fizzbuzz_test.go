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
func TestInput_3_ShouldBe_fizz(t *testing.T) {
	fb := FizzBuzz(3)
	expected := "fizz"
	if fb != expected {
		t.Errorf("Actual = %s, expected = %s", fb, expected)
	}
}

func TestInput_5_ShouldBe_buzz(t *testing.T) {
	fb := FizzBuzz(5)
	expected := "buzz"
	if fb != expected {
		t.Errorf("Actual = %s, expected = %s", fb, expected)
	}
}
func TestInput_6_ShouldBe_fizz(t *testing.T) {
	fb := FizzBuzz(6)
	expected := "fizz"
	if fb != expected {
		t.Errorf("Actual = %s, expected = %s", fb, expected)
	}
}

func TestInput_10_ShouldBe_buzz(t *testing.T) {
	fb := FizzBuzz(10)
	expected := "buzz"
	if fb != expected {
		t.Errorf("Actual = %s, expected = %s", fb, expected)
	}
}

func TestInput_15_ShouldBe_fizzbuzz(t *testing.T) {
	fb := FizzBuzz(15)
	expected := "fizzbuzz"
	if fb != expected {
		t.Errorf("Actual = %s, expected = %s", fb, expected)
	}
}