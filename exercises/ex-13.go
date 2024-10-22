package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(totalAmountOfPoints([]string{"3:1", "2:2", "0:1"}))

}

func totalAmountOfPoints(games []string) int {
	var points int
	const winPoints = 3
	const drawPoints = 1

	for _, game := range games {
		score := strings.Split(game, ":")

		if score[0] == score[1] {
			points += drawPoints
		}

		if score[0] > score[1] {
			points += winPoints
		}
	}

	return points
}
