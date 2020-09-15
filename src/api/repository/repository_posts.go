package repository

import (
	"net/http"

	"github.com/pragmatically-dev/apirest/src/api/database"
	"github.com/pragmatically-dev/apirest/src/api/models"
	"github.com/pragmatically-dev/apirest/src/api/repository/crud"
	"github.com/pragmatically-dev/apirest/src/api/responses"

	"github.com/pragmatically-dev/apirest/src/config"
)

//PostRepository es la interfaz de los metodos crud de los posts
type PostRepository interface {
	Save(models.Post) (models.Post, error)
	FindAll() ([]models.Post, error)
	//FindByID(primitive.ObjectID) (models.Post, error)
	//Update(primitive.ObjectID, models.Post) (primitive.ObjectID, error)
	//Delete(primitive.ObjectID) (bool, error)
}

//GetRepositoryPostCrud obtiene un puntero a la estructura RepositoryUsersCRUD
func GetRepositoryPostCrud(w http.ResponseWriter, r *http.Request) (*crud.RepositoryPostCRUD, error) {
	_, db, err := database.Connect() //se conecta a la base de datos ignorando el context mediante el underscore
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return nil, err
	}
	dbnamed := db.Database(config.DBNAME)        //obtiene el puntero a la estructura *mongo.Database
	repo := crud.NewRepositoryPostsCRUD(dbnamed) //se crea un nuevo repositorio que implementa los metodos de la interfaz userRepository
	return repo, nil
}
