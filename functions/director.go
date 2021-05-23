package functions

import (
	"context"

	connectionhelper "github.com/pramodshenkar/movieapp3/connectionHelper"
	model "github.com/pramodshenkar/movieapp3/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func AddDirector(director model.Director) (*mongo.InsertOneResult, error) {
	client, err := connectionhelper.GetMongoClient()
	if err != nil {
		return nil, err
	}
	collection := client.Database(connectionhelper.DB).Collection("directors")
	res, err := collection.InsertOne(context.TODO(), director)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func AddManyDirectors(list []model.Director) (*mongo.InsertManyResult, error) {
	insertableList := make([]interface{}, len(list))
	for i, v := range list {
		insertableList[i] = v
	}
	client, err := connectionhelper.GetMongoClient()
	if err != nil {
		return nil, err
	}
	collection := client.Database(connectionhelper.DB).Collection(connectionhelper.DIRECTORS)
	res, err := collection.InsertMany(context.TODO(), insertableList)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func GetDirectorsById(id int) (model.Director, error) {
	result := model.Director{}
	filter := bson.D{primitive.E{Key: "_id", Value: id}}
	client, err := connectionhelper.GetMongoClient()
	if err != nil {
		return result, err
	}
	collection := client.Database(connectionhelper.DB).Collection(connectionhelper.DIRECTORS)
	err = collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		return result, err
	}
	return result, nil
}

func GetAllDirectors() ([]model.Director, error) {
	filter := bson.D{{}}
	directors := []model.Director{}
	client, err := connectionhelper.GetMongoClient()
	if err != nil {
		return directors, err
	}
	collection := client.Database(connectionhelper.DB).Collection(connectionhelper.DIRECTORS)
	cur, findError := collection.Find(context.TODO(), filter)
	if findError != nil {
		return directors, findError
	}
	for cur.Next(context.TODO()) {
		t := model.Director{}
		err := cur.Decode(&t)
		if err != nil {
			return directors, err
		}
		directors = append(directors, t)
	}
	cur.Close(context.TODO())
	if len(directors) == 0 {
		return directors, mongo.ErrNoDocuments
	}
	return directors, nil
}

func UpdateDirector(director model.Director) (*mongo.UpdateResult, error) {
	filter := bson.D{primitive.E{Key: "_id", Value: director.ID}}

	updater := bson.D{primitive.E{Key: "$set", Value: bson.D{{Key: "name", Value: director.Name}, {Key: "address", Value: director.Address}}}}

	client, err := connectionhelper.GetMongoClient()
	if err != nil {
		return nil, err
	}
	collection := client.Database(connectionhelper.DB).Collection(connectionhelper.DIRECTORS)

	res, err := collection.UpdateOne(context.TODO(), filter, updater)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func DeleteOneDirector(id int) (*mongo.DeleteResult, error) {
	filter := bson.D{primitive.E{Key: "_id", Value: id}}
	client, err := connectionhelper.GetMongoClient()
	if err != nil {
		return nil, err
	}
	collection := client.Database(connectionhelper.DB).Collection(connectionhelper.DIRECTORS)
	res, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func DeleteAllDirectors() (*mongo.DeleteResult, error) {
	selector := bson.D{{}} // bson.D{{}} specifies 'all documents'
	client, err := connectionhelper.GetMongoClient()
	if err != nil {
		return nil, err
	}
	collection := client.Database(connectionhelper.DB).Collection(connectionhelper.DIRECTORS)
	res, err := collection.DeleteMany(context.TODO(), selector)
	if err != nil {
		return nil, err
	}
	return res, nil
}
