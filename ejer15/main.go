package main

import (
	"fmt"
	"time"
)

func main() {
	canal1 := make(chan time.Duration)
	go bucle(canal1)
	fmt.Println("llegue hasta aqui")

	msg := <-canal1 // recibimos el valor del canal | similar a async await de javascript
	// imprimir el resultado
	fmt.Println(msg)

}

func bucle(canal1 chan time.Duration) {
	inicio := time.Now()
	for i := 0; i < 1000000000000; i++ {

	}

	final := time.Now()
	canal1 <- final.Sub(inicio) // asigngimos el tiempo que tarda en ejecutarse el bucle al canal
}
