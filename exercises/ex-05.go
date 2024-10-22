package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	fmt.Println(convertAStringToANumber("123"))
	fmt.Println(convertAStringToANumber("605"))
	fmt.Println(convertAStringToANumber("1405"))
	fmt.Println(convertAStringToANumber("7"))
	fmt.Println(reflect.TypeOf(convertAStringToANumber("123")))
}

func convertAStringToANumber(toInt string) int {
	number, _ := strconv.Atoi(toInt)
	return number
}
