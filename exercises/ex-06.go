package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println(convertANumberToAString(123))
	fmt.Println(convertANumberToAString(999))
	fmt.Println(convertANumberToAString(123))
	fmt.Println(reflect.TypeOf(convertANumberToAString(123)))
}

func convertANumberToAString(toString int) string {
	/*
		number := strconv.Itoa(toString)
		return number
	*/
	return fmt.Sprintf("%d", toString)
}
