package database

import (
	"context"
	"time"

	"github.com/pragmatically-dev/apirest/src/config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//Connect devuelve una instancia de mongodb
func Connect() (context.Context, *mongo.Client, error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	//defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(
		config.DBURL,
	))
	if err != nil {
		//log.Fatal(fmt.Sprintf("\n %s", err))
		return nil, nil, err
	}

	return ctx, client, nil
}
