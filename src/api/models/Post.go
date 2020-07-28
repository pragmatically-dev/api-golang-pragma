package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Post estructura que define los atributos de un post
type Post struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Title     string             `json:"title,omitempty" bson:"title,omitempty"`
	Content   string             `json:"content,omitempty" bson:"content,omitempty"`
	Author    User               `json:"author,omitempty" bson:"author,omitempty"`
	AuthorID  primitive.ObjectID `json:"author_id,omitempty" bson:"author_id,omitempty"`
	CreatedAt time.Time          `json:"create_at,omitempty" bson:"create_at,omitempty"`
	UpdatedAt time.Time          `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}

//BeforeSave se encarga de generar el id del post y de inicializar las fechas
func (p *Post) BeforeSave() error {
	p.ID = primitive.NewObjectID()
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()

	return nil
}
