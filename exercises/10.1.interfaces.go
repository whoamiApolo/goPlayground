package main

import "fmt"

type Adulto interface {
	FalaBomDia() string
	Profissao() string
}

func BomDia(pessoa Pessoa) {
	fmt.Println(pessoa.FalaBomDia())
	fmt.Println(pessoa.Profissao())
}

type Pessoa struct {
	Nome               string
	Idade              int
	Prof               string
	TempoDeExperiencia int
}

// metodo que é atrelado a pessoa e retorna uma string
func (p Pessoa) FalaBomDia() string {
	return fmt.Sprintf("%s deseja bom dia para você!", p.Nome)
}

func (p Pessoa) Profissao() string {
	return fmt.Sprintf("%s tem %d anos como %s", p.Nome, p.TempoDeExperiencia, p.Prof)
}

func main() {
	adulto := Pessoa{
		Nome:               "Thanos",
		Idade:              50,
		Prof:               "Vilão",
		TempoDeExperiencia: 10,
	}

	fmt.Println(adulto.FalaBomDia())
	fmt.Println(adulto.Profissao())

	BomDia(adulto)
}
