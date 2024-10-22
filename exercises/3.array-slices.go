package main

import (
	"fmt"
)

func main() {
	//Arrays
	var numeros [2]int
	numeros[0] = 1
	numeros[1] = 2
	println(numeros[0], numeros[1])

	var arr [2]string
	arr[0] = "Lets"
	arr[1] = "Go"
	fmt.Println(arr[0], arr[1])

	fibonacci := [5]int{1, 1, 2, 3, 5}
	fmt.Println(fibonacci)
	fmt.Println(fibonacci[0:3])

	//Slices
	numbersPares := make([]int, 3)
	numbersPares[0] = 2
	numbersPares[1] = 4
	fmt.Println(numbersPares)
	fmt.Println(len(numbersPares))

	numbersImpares := []int{1, 3, 5}
	numbersImpares = append(numbersImpares, 7, 9, 11)
	fmt.Println(numbersImpares)

}

/*
Homogêneos - Arrays e Slices, ´possuem o mesmo tipo de elementos, mas a diferença é que os arrays têm um tamanho fixo, enquanto os slices têm um tamanho dinâmico.
Heterogêneos - Maps e Structs, possuem diferentes tipos de elementos. Maps são coleções de pares chave-valor, enquanto Structs são coleções de campos.
*/
