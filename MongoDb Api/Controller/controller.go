package controller

import (
	model "Guruprasad/Model"
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://GuruDev:jTbuy6epJuhu5bXj@cluster0.ewtvlvm.mongodb.net/?retryWrites=true&w=majority"
const dbName = "netflix"
const collectionName = "watchlist"

// MOST IMPORTANT

var collection *mongo.Collection

// connect with mongoDB

func init() {
	// client options

	// clientOptions := options.Client().ApplyURI(connectionString)

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(connectionString).SetServerAPIOptions(serverAPI)

	// connect to mongo db

	client, err := mongo.Connect(context.TODO(), opts)

	if err != nil {

		log.Fatal(err)

	}

	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	fmt.Println("MongoDB Connection success")

	// var result bson.M
	// if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Decode(&result); err != nil {
	// 	panic(err)
	// }
	// fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")

	collection = client.Database(dbName).Collection(collectionName)

	// collection instance

	fmt.Println("collection instancee is ready")

}

// mongo db helpers

// insert 1 record

func insertOneMovie(movie model.Netflix) error {
	inserted, err := collection.InsertOne(context.Background(), movie)
	if err != nil {
		return err
	}
	log.Println("Inserted one movie in db with id: ", inserted.InsertedID)
	return nil
}

func updateMovie(movieId string) {

	id, err := primitive.ObjectIDFromHex(movieId)
	if err != nil {

		log.Fatal(err)

	}

	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}

	result, err := collection.UpdateOne(context.Background(), filter, update)

	fmt.Println("Modified count : ", result.ModifiedCount)
}

// delete record

func deleteRecord(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId)

	filter := bson.M{"_id": id}
	deletecount, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {

		log.Fatal(err)

	}

	fmt.Println("The deleted count is : ", deletecount.DeletedCount)

}

// delete all record

func deleteAll() {
	filter := bson.D{{}}
	result, err := collection.DeleteMany(context.Background(), filter)
	if err != nil {

		log.Fatal(err)

	}
	fmt.Println("Deleted all the record : ", result.DeletedCount)

}

// get All movies form database

func getAll() []primitive.M {
	cursor, err := collection.Find(context.Background(), bson.D{{}})
	if err != nil {

		log.Fatal(err)

	}

	var movies []primitive.M

	for cursor.Next(context.Background()) {
		var movie bson.M
		err := cursor.Decode(&movie) // Pass a reference to movie
		if err != nil {
			log.Fatal(err)
		}
		movies = append(movies, movie)
	}

	defer cursor.Close(context.Background())

	return movies

}
