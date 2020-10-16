package main

import "testing"

func TestEchange(t *testing.T) {
	value1 := 10
	value2 := 5
	exchanged1 := value1
	exchanged2 := value2

	Exchange(&exchanged1, &exchanged2)

	if exchanged2 != value1 {
		t.Errorf("Values are not the same %d != %d", exchanged2, value1)
	}

	if exchanged1 != value2 {
		t.Errorf("Values are not the same %d != %d", exchanged1, value2)
	}
}
