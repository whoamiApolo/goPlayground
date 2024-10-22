package main

func main() {
	println(makeNegative(1))
	println(makeNegative(-5))
	println(makeNegative(0))
}

func makeNegative(x int) int {
	if x <= 0 {
		return x
	}

	return -x
}

/*
Nesta tarefa simples, você receberá um número e precisará torná-lo negativo. Mas talvez o número já seja negativo?
*/
