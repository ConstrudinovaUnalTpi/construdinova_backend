package bd

import (
	"log"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*Objeto de conexion a la BD*/
var MongoCN=ConectarBD()
var clientOptions = options.Client().ApplyURI("mongodb+srv://construdinova:abcd1234@cluster0.6uo1l.mongodb.net/construdinova_db?retryWrites=true&w=majority")

/*ConectarBD que se encraga de conectar con la BD*/
func ConectarBD() *mongo.Client {

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	log.Println("Conexion exitosa con la BD")
	return client
}

/*Chequeo connection es un ping que se le realiza a la BD*/
func ChequeoConnection() int {
	err := MongoCN.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}
	return 1
}