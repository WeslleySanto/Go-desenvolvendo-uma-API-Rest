package routes

import (
	"log"
	"net/http"

	"github.com/WeslleySanto/Go-desenvolvendo-uma-API-Rest/controllers"
	"github.com/gorilla/mux"
)

func HandleRequest() {

	r := mux.NewRouter()
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/personalidades", controllers.GetAllPersonalities).Methods("Get")
	r.HandleFunc("/api/personalidades/{id}", controllers.GetPersonalitiesById).Methods("Get")
	log.Fatal(http.ListenAndServe(":8000", r))
}
