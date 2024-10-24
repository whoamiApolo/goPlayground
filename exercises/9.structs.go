package main

import "fmt"

type Animal struct {
	Nome   string
	Idade  int
	Peso   float64
	Altura float64
	Cor    string
}

func showBasic(Nome string, Idade int, Peso float64, Altura float64, Cor string) string {
	return fmt.Sprintf("Nome: %s, Idade: %d, Peso: %.2f, Altura: %.2f, Cor: %s", Nome, Idade, Peso, Altura, Cor)
}

func showIntermediary(animal Animal) string {
	return fmt.Sprintf("Nome: %s, Idade: %d, Peso: %.2f, Altura: %.2f, Cor: %s", animal.Nome, animal.Idade, animal.Peso, animal.Altura, animal.Cor)
}

func (a Animal) showAdvanced() string {
	return fmt.Sprintf("Nome: %s, Idade: %d, Peso: %.2f, Altura: %.2f, Cor: %s", a.Nome, a.Idade, a.Peso, a.Altura, a.Cor)

}

func main() {
	animal := Animal{
		Nome:   "Cachorro",
		Idade:  2,
		Peso:   10.5,
		Altura: 0.5,
		Cor:    "Marrom",
	}

	fmt.Println(showBasic(animal.Nome, animal.Idade, animal.Peso, animal.Altura, animal.Cor))
	fmt.Println(showIntermediary(animal))
	fmt.Println(animal.showAdvanced())
}
