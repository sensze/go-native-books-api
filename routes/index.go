package routes

import "github.com/gorilla/mux"

func IndexRoute(router *mux.Router) {
	api := router.PathPrefix("/api").Subrouter()

	AuthorRoutes(api)
}