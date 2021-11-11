package main

import "fmt"

var Calculo func(int, int) int = func(num1 int, num2 int) int {
	return num1 + num2
}

func main() {
	fmt.Printf("Sumo 5 + 7 = %d \n", Calculo(5, 7))

	// Restar
	Calculo = func(num1 int, num2 int) int {
		return num1 - num2
	}

	fmt.Printf("Resto 5 - 7 = %d \n", Calculo(5, 7))

	// Dividir
	Calculo = func(num1 int, num2 int) int {
		return num1 / num2
	}

	fmt.Printf("Divido 12 / 3 = %d \n", Calculo(12, 3))

	// fmt.Println(Operaciones())

	// closures
	tablaDel := 2
	MiTabla := Tabla(tablaDel)

	for i := 1; i < 11; i++ {
		fmt.Printf("%d x %d = %d \n", tablaDel, i, MiTabla())
	}
}

func Operaciones() func() int {
	resultado := func() int {
		var a int = 23
		var b int = 14
		return a + b
	}

	return resultado
}

func Tabla(valor int) func() int {
	numero := valor
	secuencia := 0

	return func() int {
		secuencia++
		return numero * secuencia
	}
}
