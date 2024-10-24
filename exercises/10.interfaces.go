package main

import "fmt"

type Animal interface {
	Falar() string
	Comer() int
}

func DormeODiaTodo(animal Animal) {
	fmt.Println(animal.Falar())
	fmt.Println(animal.Comer())
}

type Cachorro struct {
	Raca string
	Cor  string
	Peso int
}

type Gato struct {
	Cor  string
	Peso int
}

func (c Cachorro) Falar() string {
	return "Au au"
}

func (c Cachorro) Comer() int {
	return c.Peso
}

func (g Gato) Falar() string {
	return "Miau"
}

func (g Gato) Comer() int {
	return g.Peso
}

func main() {
	cachorro := Cachorro{
		Raca: "Pitbull",
		Cor:  "Marrom",
		Peso: 10,
	}

	gato := Gato{
		Cor:  "Branco",
		Peso: 5,
	}

	DormeODiaTodo(cachorro)
	DormeODiaTodo(gato)

	/*
		fmt.Println(cachorro.Falar())
		fmt.Println(cachorro.Comer())

		fmt.Println(gato.Falar())
		fmt.Println(gato.Comer())
	*/

}

// interface é uma forma de agrupar métodos de diferentes tipos de structs
