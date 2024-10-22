package main

func main() {
	println(summation(8))
	println(summation(2))
}

func summation(values int) int {
	aux := 0
	for i := 0; i <= values; i++ {
		aux += i
	}

	return aux
}
