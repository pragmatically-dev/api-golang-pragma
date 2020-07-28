package auto

import (
	"fmt"
	"log"

	"github.com/pragmatically-dev/apirest/src/api/database"
	"github.com/pragmatically-dev/apirest/src/config"
)

//Load this function upload the data to the db
func Load() {
	ctx, db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}

	defer db.Disconnect(ctx)
	/*
		for _, user := range users {
			user.BeforeSave()
			result, err := db.Database(config.DBNAME).Collection("Users").InsertOne(ctx, &user)
			if err != nil {
				log.Fatalln(err)
			}
			fmt.Println(result)
		}
	*/
	for _, post := range posts {
		result, err := db.Database(config.DBNAME).Collection("Posts").InsertOne(ctx, &post)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println(result)
	}
}
