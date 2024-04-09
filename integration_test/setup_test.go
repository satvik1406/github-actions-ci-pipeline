package integration

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"helloworld/domain"
	"helloworld/models"
	"time"
)

const (
	queryTruncateMessage = "TRUNCATE TABLE messages;"
	queryInsertMessage  = "INSERT INTO messages(title, body, created_at) VALUES(?, ?, ?);"
)


func database() {
	domain.InitialiseMongoDB(true)
}

func refreshMessagesTable() error {

	_, err := domain.MongoCollection.DeleteMany(context.TODO(),bson.M{})
	if err != nil {
		panic(err.Error())
	}
	return nil
}

func seedOneMessage() (models.Message, error) {
	msg := models.Message{
		Id: 1,
		Title:     "the title",
		Body:      "the body",
		CreatedAt: time.Now(),
	}

	result, err := domain.MongoCollection.InsertOne(context.TODO(), msg)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Inserted document with _id: %v\n", result.InsertedID)
	return msg, nil
}
