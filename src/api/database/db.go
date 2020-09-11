<<<<<<< HEAD
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
		return nil, nil, err
	}

	return ctx, client, nil
}
=======
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
		return nil, nil, err
	}

	return ctx, client, nil
}
>>>>>>> 44d8fadc7ecbfbf78708d2b012037f3c8fd955b7
