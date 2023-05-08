package main

import (
	"fmt"

	"github.com/WeslleySanto/Go-desenvolvendo-uma-API-Rest/models"
	"github.com/WeslleySanto/Go-desenvolvendo-uma-API-Rest/routes"
)

func main() {
	models.Personalidades = []models.Personalidade{
		{"Nome 1", "Historia 1"},
		{"Nome 2", "Historia 2"},
	}

	fmt.Println("Iniciando o servidor Rest com GO")
	routes.HandleRequest()
}
