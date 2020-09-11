package models

import (
	"context"
	"errors"
	"time"

	"github.com/badoux/checkmail"

	"github.com/pragmatically-dev/apirest/src/api/security"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

//User es el modelo de usuario
type User struct {
	ID        primitive.ObjectID   `json:"_id,omitempty" bson:"_id,omitempty"`
	Nickname  string               `json:"nick_name,omitempty" bson:"nick_name,omitempty"`
	Email     string               `json:"email,omitempty" bson:"email,omitempty"`
	Password  string               `json:"password,omitempty" bson:"password,omitempty"`
	CreatedAt time.Time            `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt time.Time            `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
	Posts     []primitive.ObjectID `json:"posts,omitempty" bson:"posts,omitempty"`
} //TODO: Cambiar el tipo de Posts por un slice de String

//UpdateAt se encarga de acutalizar la fecha de al actualizar un registro
func (u *User) UpdateAt() {
	u.UpdatedAt = time.Now()
}

//Verify se encarga de comprobar si el password y el email son validos
func (u *User) Verify() (bool, error) {
	if u.Email == "" {
		return false, errors.New("El campo email no puede estar vacio")
	}
	if u.Password == "" {
		return false, errors.New("El campo password no puede estar vacio")
	}
	if len(u.Password) < 8 {
		return false, errors.New("El campo password tiene que tener una longitud mayor a 8 digitos")
	}
	if err := checkmail.ValidateFormat(u.Email); err != nil {
		return false, errors.New("El email no tiene el formato correcto")
	}
	if err := checkmail.ValidateHost(u.Email); err != nil {
		return false, errors.New("El host no responde,por favor ingrese un email valido")
	}
	return true, nil
}

//BeforeSave pasa el password por una funcion de hash
func (u *User) BeforeSave() error {
	isValid, err := u.Verify()
	if err != nil {
		return err
	}
	if isValid {
		u.ID = primitive.NewObjectID()
		u.CreatedAt = time.Now()
		u.UpdatedAt = time.Now()
		hashedPass, err := security.Hash(u.Password)
		if err != nil {
			return err
		}
		u.Password = string(hashedPass)
		return nil
	} else {
		return err
	}
}

//ValidateAvailability comprueba si el email y el nick estan disponibles para su uso
func (u *User) ValidateAvailability(collection *mongo.Collection) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	email := collection.FindOne(ctx, User{Email: u.Email})
	nick := collection.FindOne(ctx, User{Nickname: u.Nickname})
	if email.Err() != nil && nick.Err() != nil {
		return true, nil
	} else {
		return false, errors.New("El Nickname o el Email ya estan siendo utilizados")
	}
}
