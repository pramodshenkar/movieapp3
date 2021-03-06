package functions

import (
	"context"

	connectionhelper "github.com/pramodshenkar/movieapp3/connectionHelper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func ReplaceDirectorInMovie(movieid int, directorid int) (*mongo.UpdateResult, error) {
	filter := bson.D{primitive.E{Key: "_id", Value: movieid}}

	updater := bson.D{
		primitive.E{
			Key: "$set",
			Value: bson.D{
				{Key: "director", Value: directorid},
			},
		},
	}

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

func AddProducerToMovie(movieid int, producerid int) (*mongo.UpdateResult, error) {
	filter := bson.D{primitive.E{Key: "_id", Value: movieid}}

	updater := bson.D{
		primitive.E{
			Key: "$addToSet",
			Value: bson.D{
				{Key: "producers", Value: producerid},
			},
		},
	}

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

func RemoveProducerFromMovie(movieid int, producerid int) (*mongo.UpdateResult, error) {
	filter := bson.D{primitive.E{Key: "_id", Value: movieid}}

	updater := bson.D{
		primitive.E{
			Key: "$pull",
			Value: bson.D{
				{Key: "producers", Value: producerid},
			},
		},
	}

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

func AddActorToMovie(movieid int, actorid int, role string, salary string) (*mongo.UpdateResult, error) {
	filter := bson.D{primitive.E{Key: "_id", Value: movieid}}

	updater := bson.D{
		primitive.E{
			Key: "$addToSet",
			Value: bson.D{
				{
					Key: "actors",
					Value: bson.D{
						{Key: "actorid", Value: actorid},
						{Key: "role", Value: role},
						{Key: "salary", Value: salary},
					},
				},
			},
		},
	}

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

func RemoveActorFromMovie(movieid int, actorid int) (*mongo.UpdateResult, error) {
	filter := bson.D{primitive.E{Key: "_id", Value: movieid}}

	updater := bson.D{
		primitive.E{
			Key: "$pull",
			Value: bson.D{
				{
					Key: "actors",
					Value: bson.D{
						{Key: "actorid", Value: actorid},
					},
				},
			},
		},
	}

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
