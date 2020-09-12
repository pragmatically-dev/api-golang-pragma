package repository

import "github.com/pragmatically-dev/apirest/src/api/models"

//PostRepository es la interfaz de los metodos crud de los posts
type PostRepository interface {
	Save(models.Post) (models.Post, error)
	//FindAll() ([]models.Post, error)
	//FindByID(primitive.ObjectID) (models.Post, error)
	//Update(primitive.ObjectID, models.Post) (primitive.ObjectID, error)
	//Delete(primitive.ObjectID) (bool, error)
}
    