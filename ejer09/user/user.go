package user

import "time"

type Usuario struct {
	Id        int
	Nombre    string
	FechaAlta time.Time
	Status    bool
}

func (this *Usuario) AltaUsuario(id int, nombre string, FechaAlta time.Time, status bool) {
	this.Id = id
	this.Nombre = nombre
	this.FechaAlta = FechaAlta
	this.Status = status
}
