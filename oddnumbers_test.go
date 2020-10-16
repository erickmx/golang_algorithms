package main

import "testing"

func TestOddNumbers(t *testing.T) {
	nextOddNumber := OddNumbersGenerator()
	var number uint64
	for i := 0; i < 10000; i++ {
		number = nextOddNumber()
		if number%2 == 0 {
			t.Errorf("Number %d is not an odd number", number)
		}
	}
}
