package bd

import (
	"context"
	"time"

	"github.com/felipe1297/construdinova_backend/models"
	"go.mongodb.org/mongo-driver/bson"
)

func ChequeoYaExisteUsuario(email string) (models.User, bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 15 *- time.Second)
	defer cancel()

	db := MongoCN.Database("construdinova_db")
	col := db.Collection("user")

	condicion := bson.M{"email":email}

	var resultado models.User

	err := col.FindOne(ctx, condicion).Decode(&resultado)
	ID := resultado.ID.Hex()
	
	if err != nil {
		return resultado, false, ID
	}

	return resultado, true, ID
}