package models

/*Semestre modelo de usuario*/
type Semestre struct {
	Numero   string      `json:"numero"`
	Materias []*Materias `json:"materias"`
}
