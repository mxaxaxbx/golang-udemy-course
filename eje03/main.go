package main

import "fmt"

var estado bool

func main() {
	estado = true

	switch estado {
	case true:
		fmt.Println("Estado es true")

	case false:
		fmt.Println("Estado es false")

	default:
		fmt.Println("Estado es default")
	}
}
