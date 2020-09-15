package middlewares

import (
	"log"
	"net/http"
)

//SetMiddlewareLogger se encarga de mostrar todas las peticiones que van llegando al sevidor
func SetMiddlewareLogger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s%s %s", r.Method, r.Host, r.RequestURI, r.Proto)
		next(w, r)
	}
}

//SetMiddlewareJSON setea que todos las respuestas del server se manden en json
func SetMiddlewareJSON(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next(w, r)
	}
}

//SetMiddleWareAuthentication verifica que para acceder a la api se este autenticado
func SetMiddleWareAuthentication(next http.HandlerFunc) http.HandlerFunc {
	/*
		TODO: Buscar el token en la base de datos y verificar si es válido
	*/
	return func(w http.ResponseWriter, r *http.Request) {
		//auth := r.URL.Query().Get("auth")
		auth2 := r.Header.Get("Token")
		//TODO: Crear un servicio de verificación de tokens
		if auth2 == "" {
			//responses.JSON(w, http.StatusForbidden, "No estas autorizado o token no expirado")
			next(w, r)
		} else {
			next(w, r)
		}
	}
}
