package models

import (
	"errors"
	"html"
	"strings"
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
	p.Title = html.EscapeString(strings.TrimSpace(p.Title))
	p.Author = User{}
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()

	return nil
}    

//Validate verifica que los campos requeridos no esten vacios
func (p Post) Validate() error {
	if p.Title == "" {
		return errors.New("El titulo es requerido")
	}
	if p.Content == "" {
		return errors.New("El contenido es requerido")
	}
	if p.ID == primitive.NilObjectID {
		return errors.New("El autor es requerido")
	}
	return nil
}
          