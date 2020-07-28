package config

import (
	"os"
	"strconv"
)

// Exporta el port
var (
	IP     = "localhost"
	PORT   = 0
	DBURL  = ""
	DBNAME = "restapi"
	EPASS  = ""
)

//Load se encarga de configurar el puerto de escucha del servidor y las variables de entorno
func Load() {
	var err error
	PORT, err = strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		PORT = 5000
	}

	DBURL = os.Getenv("DB_URI")
	EPASS = os.Getenv("EPASS")
}
