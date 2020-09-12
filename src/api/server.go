package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pragmatically-dev/apirest/src/api/emails"
	"github.com/pragmatically-dev/apirest/src/api/router"
	"github.com/pragmatically-dev/apirest/src/config"
)

//Run start the server
func Run() {
	config.Load()
	//auto.Load()
	//emails.Test()
	fmt.Printf("\nServer on port %d", config.PORT)
	fmt.Printf("\n\nDB IS CONNECTED\n")
	Listen(config.IP, config.PORT)
}

//Listen configura el listenner
func Listen(IP string, PORT int) {
	var r *mux.Router
	r = router.New()
	port := fmt.Sprintf("%s:%d", IP, PORT)
	log.Fatal(http.ListenAndServe(port, r))

}
