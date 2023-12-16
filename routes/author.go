package routes

import (
	authorcontroller "go-native-api/controllers/author_controller"

	"github.com/gorilla/mux"
)

func AuthorRoutes(routes *mux.Router) {
	router := routes.PathPrefix("/author").Subrouter()

	router.HandleFunc("", authorcontroller.Index).Methods("GET")
	router.HandleFunc("", authorcontroller.Create).Methods("POST")
	router.HandleFunc("/detail/{id}", authorcontroller.Detail).Methods("GET")
	router.HandleFunc("/update/{id}", authorcontroller.Update).Methods("PUT")
	router.HandleFunc("/delete/{id}", authorcontroller.Delete).Methods("DELETE")
}