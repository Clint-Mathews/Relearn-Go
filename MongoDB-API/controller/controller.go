package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Clint-Mathews/ReLearnGo/MongoDbAPI/41/model"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://go:123@cluster0.w3hkncf.mongodb.net/?retryWrites=true&w=majority"
const databaseName = "netflix"
const collectionName = "watchlist"

// Imp
var collection *mongo.Collection

// connect with mongodb

func init() {
	// Client Options
	clientOptions := options.Client().ApplyURI(connectionString)

	//Connect to Mongo DB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
		log.Fatal("MongoDb connection success")
	}

	// Collection instance
	collection = client.Database(databaseName).Collection(collectionName)
	fmt.Println("Collection instance is ready")
}

// MongoDB Helpers

// Insert 1 record
func insertOneMovie(movie model.Netflix) {
	inserted, err := collection.InsertOne(context.Background(), movie)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inerted 1 movie in db with id", inserted.InsertedID)
}

func updateOneRecord(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}
	updated, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Updated ", updated.UpsertedCount, "records")
}

func deleteOneRecord(movieId string) int64 {
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}
	deleteResult, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Deleted count ", deleteResult.DeletedCount)
	return deleteResult.DeletedCount
}

func deleteAllRecords() int64 {
	deleteResult, err := collection.DeleteMany(context.Background(), bson.D{{}}, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Deleted count ", deleteResult.DeletedCount)
	return deleteResult.DeletedCount
}

// Get all movies

func getAllMovies() []primitive.M {
	cursor, err := collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
	var movies []primitive.M

	for cursor.Next(context.Background()) {
		var movie bson.M
		if err := cursor.Decode(&movie); err != nil {
			log.Fatal(err)
		}
		movies = append(movies, movie)
	}
	defer cursor.Close(context.Background())
	return movies
}

// Controller
func GetAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	allMovies := getAllMovies()
	json.NewEncoder(w).Encode(allMovies)
}

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")
	var movie model.Netflix
	_ = json.NewDecoder(r.Body).Decode(&movie)
	insertOneMovie(movie)
	json.NewEncoder(w).Encode(movie)
}

func MarkAsWatched(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "PUT")
	params := mux.Vars(r)
	updateOneRecord(params["id"])
	json.NewEncoder(w).Encode(params)
}

func DeleteOneMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")
	params := mux.Vars(r)
	deleteCount := deleteOneRecord(params["id"])
	json.NewEncoder(w).Encode(deleteCount)
}

func DeleteAllMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")
	deleteCount := deleteAllRecords()
	json.NewEncoder(w).Encode(deleteCount)
}
