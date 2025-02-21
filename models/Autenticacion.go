package models

/*Autenticacion modelo de usuario*/
type Autenticacion struct {
	DOCUMENT          string
	CODE              string
	REGISTRATION_DATE string
	STATE             int32
}
