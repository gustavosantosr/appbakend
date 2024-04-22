package models

/*Resultados modelo de Resultados*/
type Resultados struct{
	Item string
	Correcta string
	Respuesta string
	Documento string
}
/*ItemPregunta modelo de ItemPregunta*/
type ItemPregunta struct{
	Item string
	Correcta string
}
/*OpcionRespuesta modelo de OpcionRespuesta*/
type OpcionRespuesta struct{
	Opcion string
}	
/*Consolidados modelo de Consolidados*/
type Consolidados struct{
	EMPLID string
}	
/*Componentes modelo de Componentes*/
type Componentes struct{
	Componente string
}	
/*ComponenteNumeros modelo de Componentes*/
type ComponenteNumeros struct{
	Componente string
	Cuenta int
}	
/*Competencias modelo de Competencias*/
type Competencias struct{
	Componente string
	Competencia string
}	
/*CompetenciaNumeros modelo de Competencias*/
type CompetenciaNumeros struct{
	Competencia string
	Cuenta int
}	
/*AplicacionFecha modelo de Competencias*/
type AplicacionFecha struct{
	Fecha string
}	
/*Criterios modelo de Competencias*/
type Criterios struct{
	Componente string
	Competencia string
	Afirmacion string
	Evidencia string
	Tarea string
}	

/*Estudiantes modelo de Competencias*/
type Estudiantes struct{
	Nombre string
	IDEstudiante string
}	
/*RespuestaEstudiante modelo de Competencias*/
type RespuestaEstudiante struct{
	Respuesta string
}	
