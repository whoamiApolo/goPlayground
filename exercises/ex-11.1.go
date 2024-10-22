package main

import "fmt"

func main() {
	var aux bool = true

	// Criando o map para simular o operador ternário
	resultado := map[bool]string{true: "Aprovado", false: "Reprovado"}[aux]

	fmt.Println("Resultado:", resultado)
}
