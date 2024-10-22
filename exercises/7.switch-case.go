package main

func main() {
	classificacao := 3

	switch classificacao {
	case 1:
		println("Primeiro lugar")
	case 2:
		println("Segundo lugar")
	case 3:
		println("Terceiro lugar")
	default:
		println("Não subiu ao pódio")
	}

	hero := "Homem-Aranha"

	switch hero {
	case "Thanos":
		println("Thanos é o vilão")
	case "Homem-Aranha":
		println("Homem-Aranha é o herói")
	default:
		println("Personagem desconhecido")
	}
}
