package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(RepeatStr(6, "Lets go! "))
}

func RepeatStr(repeat int, value string) string {
	return strings.Repeat(value, repeat)
	/*
			for i := 0; i < 6; i++ { x += aux	}
		fmt.Println(x)
	*/
}
