package crud

import (
	"context"
	"time"

	"github.com/pragmatically-dev/apirest/src/api/utils/channels"

	"github.com/pragmatically-dev/apirest/src/api/models"
	"go.mongodb.org/mongo-driver/mongo"
)

//RepositoryPostCRUD esta estructura contiene un puntero de la estructura DATABASE
type RepositoryPostCRUD struct {
	db *mongo.Database
}

//NewRepositoryPostsCRUD se encarga de retornar una referencia de la estructura RepositoryUsersCRUD
func NewRepositoryPostsCRUD(db *mongo.Database) *RepositoryPostCRUD {
	return &RepositoryPostCRUD{db}
}

//Save implementa la interfaz PostRepository
func (repository *RepositoryPostCRUD) Save(post models.Post) (models.Post, error) {
	var globalError error
	collection := repository.db.Collection("Posts")
	//se crea un canal por donde las go routines van a comunicar los posibles errores mediante un boleano
	done := make(chan bool)
	//se define un contexto para realizar las querys
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	//se crea una go routine que se encarga de manejar las querys concurrentemente
	go func(ch chan<- bool) {
		defer close(ch)
		err := post.BeforeSave()
		if err != nil {
			globalError = err
			ch <- false
			return
		}
		_, err = collection.InsertOne(ctx, &post)
		if err != nil {
			globalError = err
			ch <- false
			return
		}
		ch <- true
		return
	}(done)
	if channels.OK(done) {
		return post, nil
	}
	return post, globalError
}
