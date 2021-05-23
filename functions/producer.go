package functions

import (
	"context"

	connectionhelper "github.com/pramodshenkar/movieapp3/connectionHelper"
	model "github.com/pramodshenkar/movieapp3/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func AddProducer(producer model.Producer) (*mongo.InsertOneResult, error) {
	client, err := connectionhelper.GetMongoClient()
	if err != nil {
		return nil, err
	}
	collection := client.Database(connectionhelper.DB).Collection(connectionhelper.PRODUCERS)
	res, err := collection.InsertOne(context.TODO(), producer)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func AddManyProducers(list []model.Producer) (*mongo.InsertManyResult, error) {
	insertableList := make([]interface{}, len(list))
	for i, v := range list {
		insertableList[i] = v
	}
	client, err := connectionhelper.GetMongoClient()
	if err != nil {
		return nil, err
	}
	collection := client.Database(connectionhelper.DB).Collection(connectionhelper.PRODUCERS)
	res, err := collection.InsertMany(context.TODO(), insertableList)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func GetProducersById(id int) (model.Producer, error) {
	result := model.Producer{}
	filter := bson.D{primitive.E{Key: "_id", Value: id}}
	client, err := connectionhelper.GetMongoClient()
	if err != nil {
		return result, err
	}
	collection := client.Database(connectionhelper.DB).Collection(connectionhelper.PRODUCERS)
	err = collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		return result, err
	}
	return result, nil
}

func GetAllProducers() ([]model.Producer, error) {
	filter := bson.D{{}}
	producers := []model.Producer{}
	client, err := connectionhelper.GetMongoClient()
	if err != nil {
		return producers, err
	}
	collection := client.Database(connectionhelper.DB).Collection(connectionhelper.PRODUCERS)
	cur, findError := collection.Find(context.TODO(), filter)
	if findError != nil {
		return producers, findError
	}
	for cur.Next(context.TODO()) {
		t := model.Producer{}
		err := cur.Decode(&t)
		if err != nil {
			return producers, err
		}
		producers = append(producers, t)
	}
	cur.Close(context.TODO())
	if len(producers) == 0 {
		return producers, mongo.ErrNoDocuments
	}
	return producers, nil
}

func UpdateProducer(producer model.Producer) (*mongo.UpdateResult, error) {
	filter := bson.D{primitive.E{Key: "_id", Value: producer.ID}}

	updater := bson.D{primitive.E{Key: "$set", Value: bson.D{{Key: "name", Value: producer.Name}, {Key: "address", Value: producer.Address}}}}

	client, err := connectionhelper.GetMongoClient()
	if err != nil {
		return nil, err
	}
	collection := client.Database(connectionhelper.DB).Collection(connectionhelper.PRODUCERS)

	res, err := collection.UpdateOne(context.TODO(), filter, updater)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func DeleteOneProducer(id int) (*mongo.DeleteResult, error) {
	filter := bson.D{primitive.E{Key: "_id", Value: id}}
	client, err := connectionhelper.GetMongoClient()
	if err != nil {
		return nil, err
	}
	collection := client.Database(connectionhelper.DB).Collection(connectionhelper.PRODUCERS)
	res, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func DeleteAllProducers() (*mongo.DeleteResult, error) {
	selector := bson.D{{}} // bson.D{{}} specifies 'all documents'
	client, err := connectionhelper.GetMongoClient()
	if err != nil {
		return nil, err
	}
	collection := client.Database(connectionhelper.DB).Collection(connectionhelper.PRODUCERS)
	res, err := collection.DeleteMany(context.TODO(), selector)
	if err != nil {
		return nil, err
	}
	return res, nil
}
