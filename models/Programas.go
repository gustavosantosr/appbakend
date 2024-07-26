package models

/*Programas modelo de usuario*/
type Programas struct {
	TipoCarrera                   string
	CodAcademicoPrograma          string
	ProgramaDescripcion           string
	ProgramaDuracion              string
	ProgramaPerfilAplicacion      string
	ProgramaPerfilProfesional     string
	ProgramaCompromisoProfesional string
	ProgramaPerfilOcupacional     string
}

/*Programas modelo de usuario*/
type ProgramasbyEmplid struct {
	Emplid              string
	TipoCarrera         string
	Status              string
	CodPlanAcademico    string
	ProgramaDescripcion string
	ProgramaDuracion    string
}
