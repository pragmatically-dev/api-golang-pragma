package repository

import "github.com/pragmatically-dev/apirest/src/api/models"

type PostRepository interface {
	Save(models.Post) (models.Post, error)
	//FindAll() ([]models.Post, error)
	//FindByID(primitive.ObjectID) (models.Post, error)
	//Update(primitive.ObjectID, models.Post) (primitive.ObjectID, error)
	//Delete(primitive.ObjectID) (bool, error)
}
