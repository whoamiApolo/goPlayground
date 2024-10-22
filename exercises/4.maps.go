package main

import "fmt"

func main() {
	books := map[string]int{}
	books["Harry Potter"] = 1
	books["The Lord of the Rings"] = 2
	fmt.Println(books)
	fmt.Println(books["Harry Potter"])
	fmt.Println(books["The Lord of the Rings"])
	//Outra forma de fazer
	movies := map[string]int{
		"Avatar":  1,
		"Titanic": 2,
	}
	fmt.Println(movies)
	fmt.Println(movies["Avatar"])
	fmt.Println(movies["Titanic"])
	movies["Velozes e Furiosos"] = 3
	fmt.Println(movies)
}
