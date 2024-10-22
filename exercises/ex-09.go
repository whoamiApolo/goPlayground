package main

import "fmt"

func main() {
	fmt.Println(EvenOrOdd(7))
}

func EvenOrOdd(number int) string {
	if number%2 == 0 {
		return "7 is even"
	}

	return "7 is odd"
}
