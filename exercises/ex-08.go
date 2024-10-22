package main

import "fmt"

func main() {
	fmt.Println(invertValues([]int{1, 2, 3, 4, 5}))
	fmt.Println(invertValues([]int{1, -2, 3, -4, 5}))
	fmt.Println(invertValues([]int{}))
}

func invertValues(numbers []int) []int {
	var aux []int
	for _, number := range numbers {
		aux = append(aux, number*(-1))
	}

	return aux
}
