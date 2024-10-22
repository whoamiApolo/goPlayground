package main

import "fmt"

func main() {
	fmt.Println(rockPaperScissors("rock", "paper"))
}

func rockPaperScissors(thanos, loki string) string {
	if thanos == loki {
		return "Draw"
	}

	if thanos == "rock" && loki == "scissors" {
		return "Thanos wins"
	}

	if thanos == "scissors" && loki == "paper" {
		return "Thanos wins"
	}

	if thanos == "paper" && loki == "rock" {
		return "Thanos wins"
	}

	return "Loki win"
}
