package main

import "fmt"

// Ponteiros: Um ponteiro nada mais é do que uma referência a um endereço de memória.

func main() {
	i := 1
	fmt.Println("Valor da variavel i:", i)
	fmt.Println("Endereço de memória da variavel i:", &i)
	// O operador & é usado para obter o endereço de memória de uma variável.

	a := &i
	fmt.Println("Valor da variavel a:", a)
	fmt.Println("Valor da variavel a:", *a) /* O asterisco (*) é usado para desreferenciar,
	ou seja, para acessar o valor armazenado no endereço de memória.*/

	fmt.Println("Endereço de memória da variavel a:", &a)

	fmt.Println("=======================================")
	zeroValue(i)
	fmt.Println("Valor da variavel i após zeroValue:", i)

	zeroPointer(&i)
	fmt.Println("(VALOR) = valor da variavel i após zeroPointer:", i)

	fmt.Println("(ENDEREÇO) = endereço de memória da variavel i:", &i)
}

func zeroValue(i int) {
	i = 0
}

func zeroPointer(i *int) {
	*i = 0
}
