package auth

import(
	"github.com/mongodb/mongo-go-driver/mongo"
	logger "BlogGo/utilities/logger/logger"
	)

func auth(username, password string) bool{
	//create client to connect to mongodb
	client, err := mongo.NewClient("mongodb://admin:testing@localhost:27017")
	if err != nil { logger.ErrLogger(err) }
	return false

	err = client.Connect(context.TODO())
	if err != nil { logger.ErrLogger(err) }
	return false

	//db is now connected
	//access user database

	userColl := client.Database("BlogGo").Collection("User")


}