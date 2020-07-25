package models

import (
	"time"

	"github.com/pragmatically-dev/apirest/src/api/security"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//User es el modelo de usuario
type User struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Nickname  string             `json:"nick_name,omitempty" bson:"nick_name,omitempty"`
	Email     string             `json:"email,omitempty" bson:"email,omitempty"`
	Password  string             `json:"password,omitempty" bson:"password,omitempty"`
	CreatedAt time.Time          `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt time.Time          `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}

//BeforeSave pasa el password por una funcion de hash
func (u *User) BeforeSave() error {
	u.ID = primitive.NewObjectID()
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
	hashedPass, err := security.Hash(u.Password)
	if err != nil {
		return err
	}

	u.Password = string(hashedPass)
	return nil
}
