package main

import "fmt"

func main() {
	paises := make(map[string]string, 5)
	paises["MX"] = "México"
	paises["AR"] = "Argentina"

	campeonato := map[string]int{
		"MX": 165,
		"AR": 10,
	}

	campeonato["RV"] = 5
	campeonato["CV"] = 166

	delete(campeonato, "RV")

	fmt.Println(campeonato)

	for equipo, puntos := range campeonato {
		fmt.Printf("el equipo %s tiene un puntuaje de: %d\n", equipo, puntos)
	}

	puntuaje, existe := campeonato["CV"]

	fmt.Printf("el puntuaje capturado es %d, y el equipo existe %t\n", puntuaje, existe)
}
