package models

import "database/sql"

/*Estudiante modelo de usuario*/
type Estudiante struct {
	Emplid          string
	TipoDocumento   string
	NumeroDocumento string
	Nombre          string
	Direccion       string
	Email           string
	Programas       []Programa
	Edad            sql.NullInt64
}

type Programa struct {
	CodigoPrograma string
	Programa       string
	Ety            sql.NullInt64
}
