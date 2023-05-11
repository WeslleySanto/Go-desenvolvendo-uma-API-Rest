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

func getIdByRequest(r *http.Request) string {
	vars := mux.Vars(r)
	return vars["id"]
}

func GetPersonalitiesById(w http.ResponseWriter, r *http.Request) {
	var personalidade []models.Personalidade
	database.DB.First(&personalidade, getIdByRequest(r))
	json.NewEncoder(w).Encode(personalidade)
}

func CreatePersonalities(w http.ResponseWriter, r *http.Request) {
	var novaPersonalidade models.Personalidade
	json.NewDecoder(r.Body).Decode(&novaPersonalidade)
	database.DB.Create(&novaPersonalidade)
	json.NewEncoder(w).Encode(novaPersonalidade)
}

func DeletePersonalities(w http.ResponseWriter, r *http.Request) {
	var personalidade models.Personalidade
	database.DB.Delete(&personalidade, getIdByRequest(r))
	json.NewEncoder(w).Encode(personalidade)
}

func EditPersonalities(w http.ResponseWriter, r *http.Request) {
	var personalidade models.Personalidade
	database.DB.First(&personalidade, getIdByRequest(r))
	json.NewDecoder(r.Body).Decode(&personalidade)
	database.DB.Save(&personalidade)
	json.NewEncoder(w).Encode(personalidade)
}
