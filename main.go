package main

import (
	"fmt"

	"github.com/WeslleySanto/Go-desenvolvendo-uma-API-Rest/database"
	"github.com/WeslleySanto/Go-desenvolvendo-uma-API-Rest/routes"
)

func main() {
	database.Conecta()
	fmt.Println("Iniciando o servidor Rest com GO")
	routes.HandleRequest()
}
