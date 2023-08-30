package utils

import (
	"fmt"
	"time"
)

func GetDate() time.Time {
	// Carregar o local de Brasília
	brasiliaLocation, err := time.LoadLocation("America/Sao_Paulo")
	if err != nil {
		fmt.Println("Erro ao carregar o fuso horário:", err)
		return time.Now()
	}

	// Obter a data e hora atual em Brasília
	timer := time.Now().In(brasiliaLocation)

	return timer

}

func ConvertDateLocation(date time.Time) time.Time {
	brasiliaLocation, err := time.LoadLocation("America/Sao_Paulo")
	if err != nil {
		fmt.Println("Erro ao carregar o fuso horário:", err)
	}

	// Obter a data e hora atual em Brasília
	timer := date.In(brasiliaLocation)

	return timer
}
