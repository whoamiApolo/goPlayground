package main

import (
	"fmt"
)

/*
	Strucs
	Forma de criar sua própria estrutura de dados. Personalizar de acordo com a necessidade.
	Podemos usar vários tipos diferentes.
*/

// Exemplo de struct - type < nome da struct > struct { <campos da estrutura> }
type Hero struct {
	Nome               string
	Idade              int
	Profissao          string
	FazAtividadeFisica bool
}

type Poderes struct {
	Hero
	Forca int
}

func main() {
	fmt.Println(Hero{Nome: "Thor", Idade: 1500, Profissao: "Deus do Trovão", FazAtividadeFisica: true})
	fmt.Println(Hero{Nome: "Hulk", Idade: 50, Profissao: "Cientista", FazAtividadeFisica: true})
	fmt.Println(Hero{Nome: "Homem de Ferro", Idade: 45, Profissao: "Empresário", FazAtividadeFisica: false})

	// Acessando campos da struct
	h1 := Hero{Nome: "Capitão América", Idade: 100, Profissao: "Soldado", FazAtividadeFisica: true}
	fmt.Println(h1.Nome)
	fmt.Println(h1.FazAtividadeFisica)

	// Mudar valores dos campos da struct
	h1.FazAtividadeFisica = false
	fmt.Println(h1.FazAtividadeFisica)

	h2 := Hero{Nome: "Pantera Negra", Idade: 40, Profissao: "Rei de Wakanda", FazAtividadeFisica: true}

	// Criando um slice de structs
	heroes := []Hero{}
	heroes = append(heroes, h1, h2)
	fmt.Println(heroes)

	// Criando um map de structs
	// variavel heroesMap do tipo map com chave string e valor Hero (valor é uma lista/array de Hero)
	heroesMap := map[string][]Hero{}
	heroesMap["Marvel"] = heroes
	fmt.Println(heroesMap)

	// variavel heroesMap1 do tipo map com chave string e valor Hero (valor é uma lista/array de Hero)
	heroesMap1 := map[string][]Hero{
		"DC": {{Nome: "Batman", Idade: 40, Profissao: "Bilionário", FazAtividadeFisica: true},
			{Nome: "Lanterna Verde", Idade: 30, Profissao: "Policial", FazAtividadeFisica: true}},
		"Boom! Studios": {{Nome: "Power Rangers", Idade: 20, Profissao: "Herói", FazAtividadeFisica: true}},
	}

	fmt.Println(heroesMap1)

	// Herdando campos de uma struct
	power := Poderes{h2, 100}
	fmt.Println(power)
}
