package crud

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/pragmatically-dev/apirest/src/api/database"
	"github.com/pragmatically-dev/apirest/src/api/utils/channels"
	"github.com/pragmatically-dev/apirest/src/config"

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
		err = post.Validate()
		if err != nil {
			globalError = err
			ch <- false
			return
		}
		_, db, err := database.Connect() //se conecta a la base de datos ignorando el context mediante el underscore
		if err != nil {
			globalError = err
			ch <- false
			return
		}
		dbnamed := db.Database(config.DBNAME) //obtiene el puntero a la estructura *mongo.Database
		repo := NewRepositoryUsersCRUD(dbnamed)
		_id, _ := primitive.ObjectIDFromHex(post.AuthorID)
		_, err = collection.InsertOne(ctx, &post)
		repo.PushPost(_id, post)
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

//FindAll se encarga de retornar todos los posts de la base de datos
func (repository *RepositoryPostCRUD) FindAll() ([]models.Post, error) {
	var posts []models.Post
	done := make(chan bool) //crea un canal que comunica valores boleanos
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	//go function
	//Se encarga de recuperar todos los posts de la base de datos, y transmite el resultado por medio de un canal boleano
	go func(ch chan<- bool) {
		defer cancel()
		cursor, err := repository.db.Collection("Posts").Find(ctx, bson.M{})
		defer cursor.Close(ctx)
		if err != nil {
			ch <- false
			return
		}
		for cursor.Next(ctx) { //for each element in the database
			var post models.Post
			cursor.Decode(&post)
			posts = append(posts, post) //lo acgrega al slice de posts
		}
		if len(posts) > 0 {
			ch <- true
		}
	}(done)

	if channels.OK(done) {
		return posts, nil
	}
	return []models.Post{}, nil
}
