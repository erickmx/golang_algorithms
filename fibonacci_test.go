package main

import "testing"

func TestNumber0(t *testing.T) {
	const expected = 0

	answer := Fibonacci(0)

	if answer != expected {
		t.Errorf("Expected number %d got %d", expected, answer)
	}
}

func TestNumber1(t *testing.T) {
	const expected = 1

	answer := Fibonacci(1)

	if answer != expected {
		t.Errorf("Expected number %d got %d", expected, answer)
	}
}

func TestNumber2(t *testing.T) {
	const expected = 1

	answer := Fibonacci(2)

	if answer != expected {
		t.Errorf("Expected number %d got %d", expected, answer)
	}
}

func TestNumberUntil10(t *testing.T) {
	expectedValues := []int64{
		0,
		1,
		1,
		2,
		3,
		5,
		8,
		13,
		21,
		34,
		55,
	}

	for idx, expected := range expectedValues {
		answer := Fibonacci(int64(idx))

		if answer != expected {
			t.Errorf("Expected number %d got %d", expected, answer)
		}
	}
}
