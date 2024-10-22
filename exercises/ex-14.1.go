package main

import "fmt"

func main() {
	olimpiadas := []string{"futebol", "natação", "atletismo", "ginástica", "vôlei"}

	for i := 1; i < len(olimpiadas); i++ {
		fmt.Print(olimpiadas[:i])
		fmt.Println(olimpiadas[i:])
	}
}
