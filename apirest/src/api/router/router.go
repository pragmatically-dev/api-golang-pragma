package router

import (
	"github.com/gorilla/mux"
	"github.com/pragmatically-dev/apirest/src/api/router/routes"
)

//New crea un enrutador de mux
func New() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	//return routes.SetupRoutes(router)
	return routes.SetupRoutesWithMiddlewares(router)
}
