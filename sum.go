package main

const UintSize = 32 << (^uint(0) >> 32 & 1) // 32 or 64

const (
	MaxInt  = 1<<(UintSize-1) - 1 // 1<<31 - 1 or 1<<63 - 1
	MinInt  = -MaxInt - 1         // -1 << 31 or -1 << 63
	MaxUint = 1<<UintSize - 1     // 1<<32 - 1 or 1<<64 - 1
)

func Sum(args ...int) int {
	var greatestNumber int = MinInt

	for _, number := range args {
		if number > greatestNumber {
			greatestNumber = number
		}
	}

	return greatestNumber
}
