package main

func main() {

	for i := 0; i < 5; i++ {
		println(i)
	}
	/*
		for {
			print("loop infinito")
			time.Sleep(5 * time.Second)
		}
	*/
	frutas := []string{"banana", "maçã", "uva", "morango"}

	for i, fruta := range frutas {
		println("Fruta:", fruta, "Indice:", i)
	}
}
