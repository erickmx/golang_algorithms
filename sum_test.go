package main

import (
	"fmt"
	"testing"
)

func TestGreatestInArray(t *testing.T) {
	var greatest int
	values := [][]int{
		{0, 10, 34, 12, 44, 100},
		{-11, -20, -10, 0},
		{-100, -42, -54, -32, -51, -11},
		{10},
	}

	greatestsValues := []int{
		100,
		0,
		-11,
		10,
	}

	for idx, groupValue := range values {
		greatest = Sum(groupValue...)
		if greatest != greatestsValues[idx] {
			fmt.Errorf("Expected value %d, got %d", greatestsValues[idx], greatest)
		}
	}
}
