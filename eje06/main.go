package main

import "fmt"

func main() {
	fmt.Println(calc(5, 46))
	fmt.Println(calc(2, 23, 45, 68))
	fmt.Println(calc(5))
	fmt.Println(calc(5, 46, 17, 24))
}

func uno(numero int) (z int) {
	z = numero * 2
	return
}

func dos(numero int) (int, bool) {
	if numero == 1 {
		return 5, false
	} else {
		return 10, true
	}
}

func calc(num ...int) int {
	total := 0

	for item, numero := range num {
		total = total + numero
		fmt.Printf("\n Item %d \n", item)
	}

	return total
}
