package models

import(
	"go.mongodb.org/mongo-driver/bson/primitive"
	
)
/*Apoyo modelo de usuario*/
type Apoyo struct{
	ID	primitive.ObjectID 		`bson:"_id,omitempty" json:"id"`
	Localidad string 			`bson:"Localidad" json:"Localidad,omitempty"`
	Dane12 string 				`bson:"Dane12" json:"Dane12,omitempty"`
	Vigencia int32 				`bson:"Vigencia" json:"Vigencia,omitempty"`
	CodLocalidad int32			`bson:"CodLocalidad" json:"CodLocalidad"`
	Institucion string			`bson:"Institucion" json:"Institucion,omitempty"`
	Sede string					`bson:"Sede" json:"Sede,omitempty"`
	Metodologia string			`bson:"Metodologia" json:"Metodologia,omitempty"`
	Jornada string				`bson:"Jornada" json:"Jornada,omitempty"`
	Area string					`bson:"Area" json:"Area,omitempty"`
	AreaMen string				`bson:"AreaMen" json:"AreaMen,omitempty"`
	Nota string					`bson:"Nota" json:"Nota,omitempty"`
	CountMenArea int32			`bson:"CountMenArea" json:"CountMenArea,omitempty"`
	Asignatura string			`bson:"Asignatura" json:"Asignatura,omitempty"`
	AsignaturaMen string		`bson:"AsignaturaMen" json:"AsignaturaMen,omitempty"`
	CountMenAsignatura int32	`bson:"CountMenAsignatura" json:"CountMenAsignatura,omitempty"`
}