package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	// leoArchivo()
	// leoArchivo2()
	// graboArchivo()
	graboArchivo2()
}

func leoArchivo() {
	archivo, err := ioutil.ReadFile("./file.txt")
	if err != nil {
		fmt.Println("error")
		fmt.Println(err)

	} else {
		fmt.Println(string(archivo))

	}
}

func leoArchivo2() {
	archivo, err := os.Open("./file.txt")
	if err != nil {
		fmt.Println("error")
		fmt.Println(err)

	} else {
		scanner := bufio.NewScanner(archivo)
		// lee cada linea
		for scanner.Scan() {
			registro := scanner.Text()
			fmt.Printf("linea - %s\n", registro)
		}
	}
	archivo.Close()
}

func graboArchivo() {
	archivo, err := os.Create("./new.txt")
	if err != nil {
		fmt.Println("error")
		fmt.Println(err)
		return
	}
	fmt.Fprintln(archivo, "Esta es una linea nueva")
	archivo.Close()
}

func graboArchivo2() {
	filename := "./new.txt"

	if Append(filename, "\nEsta es una segunda linea") == false {
		fmt.Println("error")
	}

}

func Append(archivo string, texto string) bool {
	// leer y escribir
	arch, err := os.OpenFile(archivo, os.O_WRONLY|os.O_APPEND, 0644)

	if err != nil {
		fmt.Println("error")
		fmt.Println(err)
		return false
	}

	_, err = arch.WriteString(texto)
	if err != nil {
		fmt.Println("error")
		fmt.Println(err)
		return false
	}

	arch.Close()
	return true

}
