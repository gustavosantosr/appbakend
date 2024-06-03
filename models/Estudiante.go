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
	CodigoPrograma  string
	Ety             string
	Programa        string
	Edad            sql.NullInt16
}
