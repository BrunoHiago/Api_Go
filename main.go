package main

import (
	"API_GO/src/config"
	"API_GO/src/routes"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// função principal
func main() {
	router := mux.NewRouter()

	//Iniciando a conexao com banco de dados
	err := config.InitDataBase()

	if err != nil {
		fmt.Println("Erro na conexao com o banco: ", err)
	}

	routes.SetupRouter(router)

	http.Handle("/", router)
	log.Fatal(http.ListenAndServe(":8000", router))
}
