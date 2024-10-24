package main

import "fmt"

type EsporteOlimpico interface {
	Esporte() string
}

func ShowEsporte(esporteOlimpico EsporteOlimpico) {
	fmt.Println(esporteOlimpico.Esporte())
}

type Modalidade struct {
	Nome string
}

func (m Modalidade) Esporte() string {
	return m.Nome
}

func main() {
	futebol := Modalidade{
		Nome: "Futebol",
	}

	tenis := Modalidade{
		Nome: "Tênis",
	}

	ShowEsporte(futebol)
	ShowEsporte(tenis)
}
