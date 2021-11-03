package main

import (
	"fmt"
)

func main() {
	var i int = 0

RUTINA:
	for i < 10 {
		if i == 4 {
			i = i + 2
			fmt.Println("voy a rutina sumando 2 a i")
			goto RUTINA
		}
		fmt.Printf("Valor de i: %d \n ", i)
		i++
	}
}
