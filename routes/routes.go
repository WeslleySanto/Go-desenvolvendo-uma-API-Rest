package routes

import (
	"log"
	"net/http"

	"github.com/WeslleySanto/Go-desenvolvendo-uma-API-Rest/controllers"
	"github.com/WeslleySanto/Go-desenvolvendo-uma-API-Rest/middleware"
	"github.com/gorilla/mux"
)

func HandleRequest() {

	r := mux.NewRouter()
	r.Use(middleware.ContentTypeJson)
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/personalidades", controllers.GetAllPersonalities).Methods("Get")
	r.HandleFunc("/api/personalidades/{id}", controllers.GetPersonalitiesById).Methods("Get")
	r.HandleFunc("/api/personalidades", controllers.CreatePersonalities).Methods("Post")
	r.HandleFunc("/api/personalidades/{id}", controllers.DeletePersonalities).Methods("Delete")
	r.HandleFunc("/api/personalidades/{id}", controllers.EditPersonalities).Methods("Put")
	log.Fatal(http.ListenAndServe(":8000", r))
}
