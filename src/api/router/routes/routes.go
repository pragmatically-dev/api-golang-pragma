package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pragmatically-dev/apirest/src/api/middlewares"
)

//Route define una estructura para recibir las request
type Route struct {
	URI     string
	Method  string
	Handler func(w http.ResponseWriter, r *http.Request)
}

//Load retorna un slice de Route cargado con las rutas
func Load() []Route {
	routes := usersRoutes

	return routes
}

//SetupRoutes Configura las rutas recibiendo como parametro un puntero a una estructura *mux.Router y lo carga con las routes definidas en userRoutes
func SetupRoutes(router *mux.Router) *mux.Router {
	for _, route := range Load() {
		router.HandleFunc(route.URI, route.Handler).Methods(route.Method) //crea un manejador por cada ruta definida en el slice de estructuras Route
	}
	return router
}

//SetupRoutesWithMiddlewares Se encarga de setear las funciones previas por donde van a pasar las peticiones antes de llegar al router como tal
func SetupRoutesWithMiddlewares(router *mux.Router) *mux.Router {
	for _, route := range Load() {
		router.HandleFunc(
			route.URI,
			middlewares.SetMiddlewareLogger( // se encarga de que las request pasen atravez del middleware
				middlewares.SetMiddlewareJSON(route.Handler)),
		).Methods(route.Method) //crea un manejador por cada ruta definida en el slice de estructuras Route
	}
	return router
}
