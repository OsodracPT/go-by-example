package main

func Sum(numbers []int) int {
	sum := 0
	// range returns index and the value by using _ we ignore the index
	for _, number := range numbers {
		sum += number
	}
	return sum
}
