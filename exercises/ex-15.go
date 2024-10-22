package main

import "fmt"

func main() {
	fmt.Println(oddEvenStringSort("CodeWars"))
}

func oddEvenStringSort(text string) string {
	var evenLetters, oddLetters string

	for i := 0; i < len(text); i++ {
		if i%2 == 0 {
			evenLetters += string(text[i])
			continue // para substituir o else
		}

		oddLetters += string(text[i])
	}

	return fmt.Sprintf("%s %s", evenLetters, oddLetters)
}
