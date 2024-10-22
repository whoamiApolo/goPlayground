package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(partsOfAList([]string{"fusca", "gol", "uno", "palio", "celta"}))
}

func partsOfAList(carros []string) string {
	var aux string
	for i := 1; i < len(carros); i++ {
		part := strings.Join(carros[:i], " ")
		partTwo := strings.Join(carros[i:], " ")
		carro := fmt.Sprintf("(%s, %s) ", part, partTwo)
		aux += carro
	}

	return aux
}
