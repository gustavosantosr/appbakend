package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*Autenticacion modelo de usuario*/
type Autenticacion struct {
	DOCUMENT          primitive.ObjectID
	CODE              string
	REGISTRATION_DATE string
	STATE             int32
}
