package functions

import (
	"context"

	connectionhelper "github.com/pramodshenkar/movieapp3/connectionHelper"
	model "github.com/pramodshenkar/movieapp3/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func AddMovie(movie model.Movie) (*mongo.InsertOneResult, error) {
	client, err := connectionhelper.GetMongoClient()
	if err != nil {
		return nil, err
	}
	collection := client.Database(connectionhelper.DB).Collection(connectionhelper.MOVIES)
	res, err := collection.InsertOne(context.TODO(), movie)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func AddManyMovies(list []model.Movie) (*mongo.InsertManyResult, error) {
	insertableList := make([]interface{}, len(list))
	for i, v := range list {
		insertableList[i] = v
	}
	client, err := connectionhelper.GetMongoClient()
	if err != nil {
		return nil, err
	}
	collection := client.Database(connectionhelper.DB).Collection(connectionhelper.MOVIES)
	res, err := collection.InsertMany(context.TODO(), insertableList)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func GetMoviesById(id int) (model.Movie, error) {
	result := model.Movie{}
	filter := bson.D{primitive.E{Key: "_id", Value: id}}
	client, err := connectionhelper.GetMongoClient()
	if err != nil {
		return result, err
	}
	collection := client.Database(connectionhelper.DB).Collection(connectionhelper.MOVIES)
	err = collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		return result, err
	}
	return result, nil
}

func GetAllMovies() ([]model.Movie, error) {
	filter := bson.D{{}}
	movies := []model.Movie{}
	client, err := connectionhelper.GetMongoClient()
	if err != nil {
		return movies, err
	}
	collection := client.Database(connectionhelper.DB).Collection(connectionhelper.MOVIES)
	cur, findError := collection.Find(context.TODO(), filter)
	if findError != nil {
		return movies, findError
	}
	for cur.Next(context.TODO()) {
		t := model.Movie{}
		err := cur.Decode(&t)
		if err != nil {
			return movies, err
		}
		movies = append(movies, t)
	}
	cur.Close(context.TODO())
	if len(movies) == 0 {
		return movies, mongo.ErrNoDocuments
	}
	return movies, nil
}

func UpdateMovie(movie model.Movie) (*mongo.UpdateResult, error) {
	filter := bson.D{primitive.E{Key: "_id", Value: movie.ID}}

	updater := bson.D{primitive.E{Key: "$set", Value: bson.D{{Key: "name", Value: movie.Name}, {Key: "budget", Value: movie.Director}, {Key: "director", Value: movie.Budget}, {Key: "producers", Value: movie.Producers}, {Key: "actors", Value: movie.Actors}}}}

	client, err := connectionhelper.GetMongoClient()
	if err != nil {
		return nil, err
	}
	collection := client.Database(connectionhelper.DB).Collection(connectionhelper.MOVIES)

	res, err := collection.UpdateOne(context.TODO(), filter, updater)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func DeleteOneMovie(id int) (*mongo.DeleteResult, error) {
	filter := bson.D{primitive.E{Key: "_id", Value: id}}
	client, err := connectionhelper.GetMongoClient()
	if err != nil {
		return nil, err
	}
	collection := client.Database(connectionhelper.DB).Collection(connectionhelper.MOVIES)
	res, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func DeleteAllMovies() (*mongo.DeleteResult, error) {
	selector := bson.D{{}} // bson.D{{}} specifies 'all documents'
	client, err := connectionhelper.GetMongoClient()
	if err != nil {
		return nil, err
	}
	collection := client.Database(connectionhelper.DB).Collection(connectionhelper.MOVIES)
	res, err := collection.DeleteMany(context.TODO(), selector)
	if err != nil {
		return nil, err
	}
	return res, nil
}
