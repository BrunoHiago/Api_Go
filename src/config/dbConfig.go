package config

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func InitDataBase() error {
	// Configurar as opções de conexão
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Criar um contexto
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Conectar ao servidor MongoDB
	c, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return err
	}

	client = c
	fmt.Println("Conexão com o MongoDB estabelecida com sucesso!")

	return nil
}

func GetClient() *mongo.Client {
	return client
}
