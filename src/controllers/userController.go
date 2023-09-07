package controllers

import (
	"API_GO/src/config"
	"API_GO/src/models"
	"API_GO/src/utils"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
)

type Response struct {
	Status  int           `json:"status"`
	Message string        `json:"message"`
	Count   int           `json:"count"`
	Data    []models.User `json:"data"`
}

func GetUsers(res http.ResponseWriter, req *http.Request) {
	var arrUser []models.User

	client := config.GetClient()
	UserCollection := client.Database("API_GO").Collection("User")

	cur, err := UserCollection.Find(context.Background(), bson.D{})
	if err != nil {
		http.Error(res, "Erro ao buscar usuários: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer cur.Close(context.Background())

	for cur.Next(context.Background()) {
		user := models.User{}

		if err := cur.Decode(&user); err != nil {
			http.Error(res, "Erro ao decodificar usuário: "+err.Error(), http.StatusInternalServerError)
			return
		}

		// Converta a data de criação, se necessário
		user.CreateAt = utils.ConvertDateLocation(user.CreateAt)
		arrUser = append(arrUser, user)
	}

	// Verifique se houve erros durante a iteração
	if err := cur.Err(); err != nil {
		http.Error(res, "Erro durante a iteração: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Configurar o cabeçalho Content-Type para application/json e enviar o JSON
	res.Header().Set("Content-Type", "application/json")
	response := Response{
		Status:  200,
		Message: "Search completed successfully",
		Data:    arrUser,
		Count:   len(arrUser),
	}

	err = json.NewEncoder(res).Encode(response)
	if err != nil {
		http.Error(res, "Erro ao codificar JSON", http.StatusInternalServerError)
		return
	}
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
		http.Error(res, "Erro ao salvar no Banco:"+errInser.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Println(id)
	response := Response{
		Status:  200,
		Message: "Saved user!",
	}

	res.Header().Set("Content-Type", "application/json")

	err = json.NewEncoder(res).Encode(response)
	if err != nil {
		http.Error(res, "Erro ao codificar JSON", http.StatusInternalServerError)
		return
	}

}

func GetUserName(res http.ResponseWriter, req *http.Request) {
	var arrUser []models.User
	vars := mux.Vars(req)

	userName := vars["name"]
	filter := bson.M{"name": userName}

	client := config.GetClient()
	UserCollection := client.Database("API_GO").Collection("User")

	cur, err := UserCollection.Find(context.Background(), filter)
	if err != nil {
		http.Error(res, "Erro ao buscar usuários: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer cur.Close(context.Background())

	for cur.Next(context.Background()) {
		user := models.User{}

		if err := cur.Decode(&user); err != nil {
			http.Error(res, "Erro ao decodificar usuário: "+err.Error(), http.StatusInternalServerError)
			return
		}

		// Converta a data de criação, se necessário
		user.CreateAt = utils.ConvertDateLocation(user.CreateAt)
		arrUser = append(arrUser, user)
	}

	err = cur.Err()
	// Verifique se houve erros durante a iteração
	if err != nil {
		http.Error(res, "Erro durante a iteração: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Configurar o cabeçalho Content-Type para application/json e enviar o JSON
	res.Header().Set("Content-Type", "application/json")
	response := Response{
		Status:  200,
		Message: "Search completed successfully",
		Data:    arrUser,
		Count:   len(arrUser),
	}

	err = json.NewEncoder(res).Encode(response)
	if err != nil {
		http.Error(res, "Erro ao codificar JSON", http.StatusInternalServerError)
		return
	}
}
