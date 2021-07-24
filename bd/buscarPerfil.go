package bd

import (
	"context"
	"log"
	"time"

	"github.com/felipe1297/construdinova_backend/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func BuscarPerfil(ID string) (models.User, error){
	
	ctx, cancel := context.WithTimeout(context.Background(), time.Second * 15)
	defer cancel()

	db := MongoCN.Database("construdinova_db")
	col := db.Collection("user")

	var perfil models.User
	objID , _ := primitive.ObjectIDFromHex(ID)

	condicion := bson.M{
		"_id": objID,
	}

	err := col.FindOne(ctx,condicion).Decode(&perfil)
	perfil.Password = ""

	if err != nil {
		log.Println("Registro no encontrado " + err.Error())
		return perfil, err
	}
	return perfil , nil
}