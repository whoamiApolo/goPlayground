package main

import "fmt"

// wellOfIdeas

func main() {
	fmt.Println(wellOfIdeas([]string{"bad", "bad", "good", "good", "good"}))
}

func wellOfIdeas(ideas []string) string {
	var goodIdeas int

	for _, idea := range ideas {
		if idea == "good" {
			goodIdeas++
		}
	}

	if goodIdeas == 0 {
		return "Fail"
	}

	if goodIdeas <= 2 {
		return "Publish!"
	}

	return "I smell a series!"
}
