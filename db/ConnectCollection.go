package db

import "go.mongodb.org/mongo-driver/mongo"

//ConnectCollection will connect to Mongodb collection that is specified
func ConnectCollection(client *mongo.Client, collName string) *mongo.Collection {
	//connect to db and specified collection
	collection := client.Database("pickMyPlan").Collection(collName)

	//return pointer to mongo.Collection type
	return collection
}
