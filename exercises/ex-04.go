package main

import "fmt"

func main() {
	fmt.Print(CountPositivesSumNegatives([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, -11, -12, -13, -14, -15}))
}

func CountPositivesSumNegatives(numbers []int) []int {
	countNumbersPositive := 0
	sumNumbersNegatives := 0

	for _, number := range numbers {
		if number > 0 {
			countNumbersPositive++
		}
		if number < 0 {
			sumNumbersNegatives += number
		}
	}

	return []int{countNumbersPositive, sumNumbersNegatives}
}
