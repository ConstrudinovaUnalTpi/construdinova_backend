package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*Modelo del de usuario de la DB*/
type User struct {
	ID			primitive.ObjectID 		`bson:"_id, omitempty" json:"id"`
	Email		string 					`bson:"email" json:"email"`
	Password	string 					`bson:"password" json:"password,omitempty"`
	TypeUser	string 					`bson:"typeUser" json:"typeUser,omitempty"`
}