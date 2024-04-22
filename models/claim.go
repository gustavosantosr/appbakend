package models

import(
	jwt "github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
)
/*Claim estructura de nuestro claim de la llave*/
type Claim struct{
	Email string `json:"email"`
	ID primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	jwt.StandardClaims
}