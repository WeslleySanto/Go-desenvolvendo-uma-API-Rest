package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/WeslleySanto/Go-desenvolvendo-uma-API-Rest/database"
	"github.com/WeslleySanto/Go-desenvolvendo-uma-API-Rest/models"
	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}

func GetAllPersonalities(w http.ResponseWriter, r *http.Request) {
	var personalidades []models.Personalidade
	database.DB.Find(&personalidades)
	json.NewEncoder(w).Encode(personalidades)
}

func GetPersonalitiesById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]

	var personalidade []models.Personalidade

	database.DB.First(&personalidade, id)

	json.NewEncoder(w).Encode(personalidade)
}
