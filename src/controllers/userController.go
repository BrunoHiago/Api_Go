package controllers

import (
	"API_GO/src/config"
	"API_GO/src/models"
	"API_GO/src/utils"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
)

func GetUsers(res http.ResponseWriter, req *http.Request) {

	var arrUser []models.User

	client := config.GetClient()
	UserCollection := client.Database("API_GO").Collection("User")

	cur, err := UserCollection.Find(context.Background(), bson.D{})

	if err != nil {
		log.Fatal(err)
	}

	for cur.Next(context.Background()) {
		user := models.User{}

		cur.Decode(&user)

		user.CreateAt = utils.ConvertDateLocation(user.CreateAt)
		arrUser = append(arrUser, user)
	}

	// Serializar o slice em JSON
	jsonData, err := json.Marshal(arrUser)
	if err != nil {
		http.Error(res, "Erro ao serializar JSON", http.StatusInternalServerError)
		return
	}

	// Configurar o cabeçalho Content-Type para application/json
	res.Header().Set("Content-Type", "application/json")

	// Escrever a resposta JSON no http.ResponseWriter
	res.Write(jsonData)
}

func CreateUser(res http.ResponseWriter, req *http.Request) {

	user := models.User{}

	err := json.NewDecoder(req.Body).Decode(&user)
	if err != nil {
		http.Error(res, "Erro ao decodificar JSON", http.StatusBadRequest)
		return
	}

	user.CreateAt = utils.GetDate()

	client := config.GetClient()
	UserCollection := client.Database("API_GO").Collection("User")

	id, errInser := UserCollection.InsertOne(context.Background(), user)

	if errInser != nil {
		http.Error(res, "Erro ao salvar no Banco:"+errInser.Error(), http.StatusBadRequest)
	}

	fmt.Println(id)

	res.WriteHeader(http.StatusOK)
}
