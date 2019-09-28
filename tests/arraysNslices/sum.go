package main

func Sum(numbers []int) int {
	sum := 0
	// range returns index and the value by using _ we ignore the index
	for _, number := range numbers {
		sum += number
	}
	return sum
}

// SumAll ... returns a sum of arrays
func SumAll(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}

	return sums
}

func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		}
	}

	return sums
}
