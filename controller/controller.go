package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/hiteshchoudhary/mongoapi/model"
	"github.com/shaurya019/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://learncodeonline:hitesh@cluster0.humov.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"
const dbName = "netflix"
const colName = "watchlist"

var collection *mongo.Collection


// connect with monogoDB
func init(){

	clientOption := options.Client().ApplyURI(connectionString)

	client, err := mongo.Connect(context.TODO(), clientOption)

	if err != nil {
		log.Fatal(err)
	}

	collection = client.Database(dbName).Collection(colName);

}


// MONGODB helpers - file

func insertOneMovie(movie model.Netflix){
	inserted,err := collection.InsertOne(context.Background(),movie);

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted 1 movie in db with id: ", inserted.InsertedID)

}


func updateOneMovie(movieId string){
	id,_ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id":id}
	update :=bson.M{"$set": bson.M{"watched": true}}

	res,err := collection.UpdateOne(context.Background(),filter,update)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("modified count: ", res.ModifiedCount)

}


func deleteOneMovie(movieId string){
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}
	deleteCount, err := collection.DeleteOne(context.Background(), filter)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("MOvie got delete with delete count: ", deleteCount)
}


func deleteAllMovie() int64{

	deleteResult,err := collection.DeleteMany(context.Background(),bson.D{{}}, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("NUmber of movies delete: ", deleteResult.DeletedCount)
	return deleteResult.DeletedCount
}


func getAllMovies() []primitive.M{
	cur,err := collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}

	var movies []primitive.M

	for cur.Next(context.Background()) {
		var movie bson.M
		err := cur.Decode(&movie)
		if err != nil {
			log.Fatal(err)
		}
		movies = append(movies, movie)
	}

	defer cur.Close(context.Background())
	return movies
}


// Actual controller - file

func GetMyAllMovies(w http.ResponseWriter, r *http.Request){}


func CreateMovie(w http.ResponseWriter, r *http.Request){}


func MarkAsWatched(w http.ResponseWriter, r *http.Request){}


func DeleteAMovie(w http.ResponseWriter, r *http.Request){}


func DeleteAllMovies(w http.ResponseWriter, r *http.Request){}