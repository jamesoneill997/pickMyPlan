package db

import "go.mongodb.org/mongo-driver/mongo"

func ConnectCollection(client *mongo.Client, collName string) *mongo.Collection {
	//connect to users db
	collection := client.Database("pickMyPlan").Collection(collName)
	return collection
}
