package main

import "log"

func main() {
	// archivo := "prueba.txt"
	// f, err := os.Open(archivo)

	// defer f.Close() // se ejecuta al final siempre

	// if err != nil {
	// 	fmt.Println("Error al abrir el archivo")
	// 	os.Exit(1)
	// }

	ejemploPanic()
}

func ejemploPanic() {
	defer func() {
		reco := recover()
		if reco != nil {
			log.Fatalf("ocurrio un error que genero panic \n %v", reco) // %v valor variable
		}
	}()

	a := 1
	if a == 1 {
		panic("Se encontró el valor de 1")
	}
}
