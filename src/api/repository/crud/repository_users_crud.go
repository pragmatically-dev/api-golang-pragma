package crud

import (
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/pragmatically-dev/apirest/src/api/models"
	"github.com/pragmatically-dev/apirest/src/api/utils/channels"
	"go.mongodb.org/mongo-driver/mongo"
)

//RepositoryUsersCRUD ASFAS
type RepositoryUsersCRUD struct {
	db *mongo.Database
}

//NewRepositoryUsersCRUD  es como un costructor
func NewRepositoryUsersCRUD(db *mongo.Database) *RepositoryUsersCRUD {
	return &RepositoryUsersCRUD{db: db}
}

//Save guarda el usuario en la base de datos
func (repository *RepositoryUsersCRUD) Save(user models.User) (models.User, error) {
	var err error
	done := make(chan bool) //crea un canal que comunica valores boleanos
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	//go function
	//Se encarga de insertar el usuario en la base de datos, y transmite el resultado por medio de un canal boleano
	go func(ch chan<- bool) {
		nick := repository.db.Collection("Users").FindOne(ctx, models.User{Nickname: user.Nickname})
		email := repository.db.Collection("Users").FindOne(ctx, models.User{Email: user.Email})
		if nick.Err() != nil && email.Err() != nil {
			var result *mongo.InsertOneResult
			user.BeforeSave()
			result, err = repository.db.Collection("Users").InsertOne(ctx, &user)
			if err != nil {
				ch <- false
				return
			}
			if result != nil {
				ch <- true
			}
		} else {
			ch <- false
			return
		}

	}(done)

	if channels.OK(done) {
		return user, nil
	} else {
		return user, errors.New("El nickname o email ya estan siendo utilizado")
	}

}

//FindAll se encarga de retornar todos los usuarios de la base de datos
func (repository *RepositoryUsersCRUD) FindAll() ([]models.User, error) {
	var users []models.User
	done := make(chan bool) //crea un canal que comunica valores boleanos
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	//go function
	//Se encarga de recuperar todos los usuarios de la base de datos, y transmite el resultado por medio de un canal boleano
	go func(ch chan<- bool) {
		defer cancel()
		cursor, err := repository.db.Collection("Users").Find(ctx, bson.M{})
		defer cursor.Close(ctx)
		if err != nil {
			ch <- false
			return
		}
		for cursor.Next(ctx) { //for each element in the database
			var user models.User
			cursor.Decode(&user)
			users = append(users, user) //lo acgrega al slice de usuarios
		}
		if len(users) > 0 {
			ch <- true
		}
	}(done)

	if channels.OK(done) {
		return users, nil
	}
	return []models.User{}, nil
}

//FindByID se encarga de retornar un elemento especifico de la base de datos
func (repository *RepositoryUsersCRUD) FindByID(_ID primitive.ObjectID) (models.User, error) {

	var user models.User
	done := make(chan bool) //crea un canal que comunica valores boleanos
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	//go function
	//Se encarga de recuperar todos los usuarios de la base de datos, y transmite el resultado por medio de un canal boleano
	go func(ch chan<- bool) {
		err := repository.db.Collection("Users").FindOne(ctx, models.User{ID: _ID}).Decode(&user)
		if err != nil {
			ch <- false
		} else {
			ch <- true
			return
		}

	}(done)

	if channels.OK(done) {
		return user, nil
	}
	return models.User{}, errors.New("No se ha encontrado el registro")
}

//Update se encarga de actualizar un registro mediante su id
func (repository *RepositoryUsersCRUD) Update(_ID primitive.ObjectID, user models.User) (primitive.ObjectID, error) {
	var userID primitive.ObjectID
	done := make(chan bool)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	go func(ch chan<- bool) {
		var Newuser = models.User{
			ID:        _ID,
			Nickname:  user.Nickname,
			Email:     user.Email,
			Password:  user.Password,
			CreatedAt: user.CreatedAt,
			UpdatedAt: time.Now(),
		}
		_, err := repository.db.Collection("Users").UpdateOne(ctx, bson.M{"_id": _ID}, bson.M{"$set": &Newuser})
		if err != nil {
			ch <- false
			return
		}
		userID = _ID
		ch <- true
		return
	}(done)

	if channels.OK(done) {
		return userID, nil
	}
	return primitive.NilObjectID, errors.New("No se ha podido actualizar el registro")
}

//TODO: Implementar Delete(primitive.ObjectID) (primitive.ObjectID, error)
