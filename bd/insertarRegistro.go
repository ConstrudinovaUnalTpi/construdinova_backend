package bd

import (
	"context"
	"time"

	"github.com/felipe1297/construdinova_backend/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)


func InsertarRegistro(u models.User) (string,bool,error){

	/*15 segundos para TimeOut*/
	ctx, cancel := context.WithTimeout(context.Background(), 15 * time.Second)
	defer cancel()

	db := MongoCN.Database("construdinova_db")
	col := db.Collection("user")

	u.Password, _ = EncriptarPassword(u.Password)

	result, err := col.InsertOne(ctx,u)

	if err != nil {
		return "", false, err
	}

	objID, _ := result.InsertedID.(primitive.ObjectID)

	return objID.String(), true, nil
}