package domain

import (
	"context"
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"helloworld/models"
	error_utils "helloworld/utils"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	MongoCollection *mongo.Collection
)

func connect(uri string)(*mongo.Client, context.Context, context.CancelFunc, error) {

	// ctx will be used to set deadline for process, here
	// deadline will of 30 seconds.
	ctx, cancel := context.WithTimeout(context.Background(), 30 * time.Second)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	return client, ctx, cancel, err
}

func ping(client *mongo.Client, ctx context.Context) error{
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		return err
	}
	fmt.Println("Connected to MongoDB successfully")
	return nil
}

func InitialiseMongoDB(test bool) {
	mongosvc := os.Getenv("MONGO_SVC")
	if mongosvc == "" {
		mongosvc = "localhost:27017"
	}
	client, ctx, _, err := connect("mongodb://"+mongosvc)
	if err != nil {
		panic(err)
	}
	//defer close(client, ctx, cancel)
	err = ping(client, ctx)
	if err != nil {
		panic(err)
	}
	if test {
		MongoCollection = client.Database("test").Collection("messages_test")
	} else {
		MongoCollection = client.Database("test").Collection("messages")
	}
}

func Create(msg *models.Message) error_utils.MessageErr {

	result, err := MongoCollection.InsertOne(context.TODO(), msg)
	if err != nil {
		return error_utils.NewInternalServerError(fmt.Sprintf("error when trying to save message: %s", err.Error()))
	}
	fmt.Printf("Inserted document with _id: %v\n", result.InsertedID)
	return  nil
}

func Update(msg *models.Message) error_utils.MessageErr {
	var res models.Message
	err := MongoCollection.FindOne(context.TODO(),bson.M{"_id":msg.Id}).Decode(&res)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return error_utils.NewNotFoundError("no record matching given id")
		}
		return error_utils.NewInternalServerError(fmt.Sprintf("error when trying to find message: %s", err.Error()))
	}
	msg.CreatedAt =  res.CreatedAt
	update := bson.M{"$set": msg}
	result, err := MongoCollection.UpdateOne(context.TODO(), bson.M{"_id":msg.Id},update)
	if err != nil {
		return error_utils.NewInternalServerError(fmt.Sprintf("error when trying to update message: %s", err.Error()))
	}
	fmt.Printf("No. of documents updated: %v\n", result.ModifiedCount)
	return  nil
}

func Get() ([]byte, error_utils.MessageErr) {
	var res models.Message
	opts := options.FindOne().SetSort(bson.D{{"created_at", -1}})
	err := MongoCollection.FindOne(context.TODO(),bson.D{},opts).Decode(&res)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, error_utils.NewNotFoundError("No record found")
		}
		return nil, error_utils.NewInternalServerError(fmt.Sprintf("error when trying to find message: %s", err.Error()))
	}

	b, err := json.Marshal(res)
	if err != nil {
		fmt.Println(err)
		return nil, nil
	}
	fmt.Printf("No. of documents updated: %v\n", string(b))
	return b,nil
}