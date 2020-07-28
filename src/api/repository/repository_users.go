package repository

import (
	"net/http"

	"github.com/pragmatically-dev/apirest/src/api/database"
	"github.com/pragmatically-dev/apirest/src/api/models"
	"github.com/pragmatically-dev/apirest/src/api/repository/crud"
	"github.com/pragmatically-dev/apirest/src/api/responses"
	"github.com/pragmatically-dev/apirest/src/config"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//UserRepository es una interfaz que especifica los metodos que van a tener en comun las estructuras que la implementan
type UserRepository interface {
	Save(models.User) (models.User, error)
	FindAll() ([]models.User, error)
	FindByID(primitive.ObjectID) (models.User, error)
	Update(primitive.ObjectID, models.User) (primitive.ObjectID, error)
	Delete(primitive.ObjectID) (bool, error)
}

//GetRepositoryCrud obtiene un puntero a la estructura RepositoryUsersCRUD
func GetRepositoryCrud(w http.ResponseWriter, r *http.Request) (*crud.RepositoryUsersCRUD, error) {
	_, db, err := database.Connect() //se conecta a la base de datos ignorando el context mediante el underscore
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return nil, err
	}
	dbnamed := db.Database(config.DBNAME)        //obtiene el puntero a la estructura *mongo.Database
	repo := crud.NewRepositoryUsersCRUD(dbnamed) //se crea un nuevo repositorio que implementa los metodos de la interfaz userRepository
	return repo, nil
}
