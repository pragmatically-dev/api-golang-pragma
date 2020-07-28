package auto

import (
	"time"

	"github.com/pragmatically-dev/apirest/src/api/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var users = []models.User{
	models.User{ID: primitive.NewObjectID(), Nickname: "santiago", Email: "santiago@gmail.com", Password: "santi1234", CreatedAt: time.Now()},
}

var user1 models.User = models.User{ID: primitive.NewObjectID(), Nickname: "santiago", Email: "santiago@gmail.com", Password: "santi1234", CreatedAt: time.Now()}

var posts = []models.Post{
	models.Post{
		ID:        primitive.NewObjectID(),
		Title:     "test post",
		Content:   "lorem ipsu",
		Author:    user1,
		AuthorID:  user1.ID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
}
