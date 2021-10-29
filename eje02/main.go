package main

import (
	"fmt"
	"strconv"
)

var Numero int
var Texto string
var Status bool = true

func main() {
	numero2, numero3, numero4, texto, status := 4, 5, 6, "texto", false
	var moneda float32 = 0
	numero2 = int(moneda)
	texto = strconv.Itoa(int(moneda))

	fmt.Println(numero2, numero3, numero4, texto)
	fmt.Println(status)

	MostrarStatus()
}

func MostrarStatus() {
	fmt.Println(Status)
}
