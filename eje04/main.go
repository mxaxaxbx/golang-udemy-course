package main

import (
	"bufio"
	"fmt"
	"os"
)

var num1 int
var num2 int
var resultado int
var leyenda string

func main() {
	fmt.Println("Ingrese numero 1:")
	fmt.Scanln(&num1)

	fmt.Println("Ingrese numero 2:")
	fmt.Scanln(&num2)

	fmt.Println("Descripcion:")

	scanner := bufio.NewScanner(os.Stdin)

	if scanner.Scan() {
		leyenda = scanner.Text()
	}

	resultado = num1 + num2

	fmt.Println(leyenda, resultado)
}
