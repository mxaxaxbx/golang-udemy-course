package main

import "fmt"

func main() {
	matriz := []int{2, 5, 4}
	fmt.Println(matriz)

	variante2()
	variante3()
	variante4()
}

func variante2() {
	elementos := [5]int{1, 2, 3, 4, 5}
	porcion := elementos[1:3]
	fmt.Println(porcion)
}

func variante3() {
	elementos := make([]int, 5, 20)
	// porcion := elementos[:3]
	fmt.Printf("Largo %d, Capacidad %d", len(elementos), cap(elementos))
}

func variante4() {
	fmt.Println("")
	nums := make([]int, 0, 0)

	for i := 0; i < 150; i++ {
		nums = append(nums, i)
	}

	fmt.Printf("Largo %d, Capacidad %d", len(nums), cap(nums))
}
