package main

import (
	"calculator/math"
	"fmt"
)

func main() {
	sum, sub, multi, div := math.Calc(5, 5)
	fmt.Println("Sum: ", sum, "\nSub: ", sub, "\nMulti: ", multi, "\nDiv: ", div)

	//e := echo.New()
}

/*
add pacotes internos
go mod init calculator

// add pacote externo
go mod tidy
*/
