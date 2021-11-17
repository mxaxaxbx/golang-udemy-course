package main

import "fmt"

type servivo interface {
	estaVivo() bool
}

type humano interface {
	respirar()
	pensar()
	comer()
	sexo() string
	estaVivo() bool
}

type animal interface {
	respirar()
	comer()
	EsCarnivoro() bool
	estaVivo() bool
}

type vegetal interface {
	ClasiVegetal() string
}

// genero humano
type hombre struct {
	edad       int
	altura     float32
	peso       float32
	respirando bool
	pensando   bool
	comiendo   bool
	esHombre   bool
	vivo       bool
}

type mujer struct {
	hombre
}

func (h *hombre) respirar() { h.respirando = true }
func (h *hombre) comer()    { h.comiendo = true }
func (h *hombre) pensar()   { h.pensando = true }
func (h *hombre) sexo() string {
	if h.esHombre {
		return "Hombre"
	} else {
		return "Mujer"
	}
}
func (h *hombre) estaVivo() bool { return h.vivo }

func HumanosRespirando(hu humano) {
	hu.respirar()
	fmt.Printf("soy un/a %s, y ya estoy respirando \n", hu.sexo())
}

// animales
type perro struct {
	respirando bool
	comiendo   bool
	carnivoro  bool
	vivo       bool
}

func (p *perro) respirar()         { p.respirando = true }
func (p *perro) comer()            { p.comiendo = true }
func (p *perro) EsCarnivoro() bool { return p.carnivoro }
func (p *perro) estaVivo() bool    { return p.vivo }

func AnimalesRespirando(an animal) {
	an.respirar()
	fmt.Printf("soy un animal, y ya estoy respirando ")
}

func AnimalesCarnivoros(an animal) int {
	if an.EsCarnivoro() {
		return 1
	} else {
		return 0
	}
}

// ser vivo
func estoyVivo(v servivo) bool {
	return v.estaVivo()
}

func main() {
	// Pedro := new(hombre)
	// Pedro.esHombre = true
	// HumanosRespirando(Pedro)

	// Maria := new(mujer)
	// HumanosRespirando(Maria)

	totalCarnivoros := 0

	Dogo := new(perro)
	Dogo.carnivoro = true
	Dogo.vivo = true
	AnimalesRespirando(Dogo)
	totalCarnivoros += AnimalesCarnivoros(Dogo)
	fmt.Println("")

	fmt.Printf("Total carnivoros: %d \n", totalCarnivoros)
	fmt.Printf("Esta vivo: %t \n", estoyVivo(Dogo))

}
