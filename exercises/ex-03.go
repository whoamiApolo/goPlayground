package main

import "fmt"

func main() {
	fmt.Println(sumOfPositive([]int{1, -4, 7, 12}))
}

func sumOfPositive(numbers []int) int {
	/*
		numbers := []int{1, -4, 7, 12} dividindo a declaração e inserindo como paremetro
		sumOfPositive(numbers []int) int
		sumOfPositive([]int{1, -4, 7, 12})
	*/
	sum := 0
	for _, number := range numbers {
		if number > 0 {
			sum += number
		}
	}
	return sum
}
