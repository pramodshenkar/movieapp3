package functions

import (
	"context"

	connectionhelper "github.com/pramodshenkar/movieapp3/connectionHelper"
	model "github.com/pramodshenkar/movieapp3/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func AddActor(actor model.Actor) (*mongo.InsertOneResult, error) {
	client, err := connectionhelper.GetMongoClient()
	if err != nil {
		return nil, err
	}
	collection := client.Database(connectionhelper.DB).Collection(connectionhelper.ACTORS)
	res, err := collection.InsertOne(context.TODO(), actor)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func AddManyActors(list []model.Actor) (*mongo.InsertManyResult, error) {
	insertableList := make([]interface{}, len(list))
	for i, v := range list {
		insertableList[i] = v
	}
	client, err := connectionhelper.GetMongoClient()
	if err != nil {
		return nil, err
	}
	collection := client.Database(connectionhelper.DB).Collection(connectionhelper.ACTORS)
	res, err := collection.InsertMany(context.TODO(), insertableList)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func GetActorsById(id int) (model.Actor, error) {
	result := model.Actor{}
	filter := bson.D{primitive.E{Key: "_id", Value: id}}
	client, err := connectionhelper.GetMongoClient()
	if err != nil {
		return result, err
	}
	collection := client.Database(connectionhelper.DB).Collection(connectionhelper.ACTORS)
	err = collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		return result, err
	}
	return result, nil
}

func GetAllActors() ([]model.Actor, error) {
	filter := bson.D{{}}
	actors := []model.Actor{}
	client, err := connectionhelper.GetMongoClient()
	if err != nil {
		return actors, err
	}
	collection := client.Database(connectionhelper.DB).Collection(connectionhelper.ACTORS)
	cur, findError := collection.Find(context.TODO(), filter)
	if findError != nil {
		return actors, findError
	}
	for cur.Next(context.TODO()) {
		t := model.Actor{}
		err := cur.Decode(&t)
		if err != nil {
			return actors, err
		}
		actors = append(actors, t)
	}
	cur.Close(context.TODO())
	if len(actors) == 0 {
		return actors, mongo.ErrNoDocuments
	}
	return actors, nil
}

func UpdateActor(actor model.Actor) (*mongo.UpdateResult, error) {
	filter := bson.D{primitive.E{Key: "_id", Value: actor.ID}}

	updater := bson.D{primitive.E{Key: "$set", Value: bson.D{{Key: "name", Value: actor.Name}, {Key: "address", Value: actor.Address}}}}

	client, err := connectionhelper.GetMongoClient()
	if err != nil {
		return nil, err
	}
	collection := client.Database(connectionhelper.DB).Collection(connectionhelper.ACTORS)

	res, err := collection.UpdateOne(context.TODO(), filter, updater)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func DeleteOneActor(id int) (*mongo.DeleteResult, error) {
	filter := bson.D{primitive.E{Key: "_id", Value: id}}
	client, err := connectionhelper.GetMongoClient()
	if err != nil {
		return nil, err
	}
	collection := client.Database(connectionhelper.DB).Collection(connectionhelper.ACTORS)
	res, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func DeleteAllActors() (*mongo.DeleteResult, error) {
	selector := bson.D{{}} // bson.D{{}} specifies 'all documents'
	client, err := connectionhelper.GetMongoClient()
	if err != nil {
		return nil, err
	}
	collection := client.Database(connectionhelper.DB).Collection(connectionhelper.ACTORS)
	res, err := collection.DeleteMany(context.TODO(), selector)
	if err != nil {
		return nil, err
	}
	return res, nil
}
