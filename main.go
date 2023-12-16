package main

import (
	"fmt"
	"go-native-api/config"
	"go-native-api/routes"
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/gorilla/mux"
)

func main(){
	config.LoadConfig()
	config.ConnectDB()
	route := mux.NewRouter()
	routes.IndexRoute(route)

	log.Println("Server running on port", config.ENV.PORT)
	http.ListenAndServe(fmt.Sprintf(":%v", config.ENV.PORT), route)
}