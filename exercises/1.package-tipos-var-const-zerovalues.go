package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Olá, Go")

	const name, age = "GoLang", 2009
	fmt.Println(name, "is", age, "years old.")

	//Pacotes
	fmt.Println(strings.Split("Lets Go", ""))

	//Tipos
	fmt.Printf("Type: %T\nValue: %v\n", true, true)
	fmt.Printf("Type: %T\nValue: %v\n", "hello go", "hello go")
	fmt.Printf("Type: %T\nValue: %v\n", 1, 2)
	fmt.Printf("Type: %T\nValue: %v\n", 1.0, 2.0)

	//Variáveis
	var sum int = 1 + 2
	fmt.Println(sum)
	num := 1
	fmt.Println(num)

	const pi = 3.1415
	fmt.Println(pi)

	//Zero Values
	var i int
	var f float64
	var b bool
	var s string

	fmt.Printf("Inteiro: %v Float: %v Booleano: %v String: %q\n", i, f, b, s)
}
