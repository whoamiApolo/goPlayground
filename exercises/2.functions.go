package main

import (
	"fmt"
)

func main() {
	valuesSum := sum(1, 2)
	fmt.Println(valuesSum)

	valuesSubtract := subtract(2, 1)
	fmt.Println(valuesSubtract)

	yourName, yourNameTwo := showName("Go")
	fmt.Println(yourName, yourNameTwo)

	firstName, lastName := fullName("Go", "Lang")
	fmt.Println(firstName, lastName)
}

func sum(x int, y int) int {
	return x + y
}

func subtract(x int, y int) int {
	return x - y
}

func showName(name string) (string, string) {
	return name, name
}

func fullName(firstName string, lastName string) (string, string) {
	return firstName, lastName
}

/*
Funções começando com minúsculas são privadas e só podem ser acessadas dentro do mesmo pacote. Por exemplo:
func sum(x int, y int) int {
	return x + y
}

Funções começando com maiúsculas são públicas e podem ser acessadas de qualquer pacote. Por exemplo:
func Sum(x int, y int) int {
	return x + y
}

Para utilizar a função Sum em outro pacote, você deve importar o pacote que contém a função. Por exemplo:
main.Sum(1, 2)
*/
