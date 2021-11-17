package main

import (
	"fmt"
	"time"

	us "./user"
)

type pepe struct {
	us.Usuario
}

func main() {
	u := new(pepe)
	u.AltaUsuario(10, "Miranda", time.Now(), true)

	fmt.Println(u.Usuario)
}
