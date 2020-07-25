package auto

import (
	"time"

	"github.com/pragmatically-dev/apirest/src/api/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var users = []models.User{
	models.User{ID: primitive.NewObjectID(), Nickname: "santiago", Email: "santiago@gmail.com", Password: "santi1234", CreatedAt: time.Now()},
}
