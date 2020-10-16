package main

func OddNumbersGenerator() func() uint64 {
	initialOdd := uint64(1)

	return func() uint64 {
		oddNumber := initialOdd
		initialOdd += 2

		return oddNumber
	}
}
