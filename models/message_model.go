package models

import (
	error_utils "helloworld/utils"
	"strings"
	"time"
)

type Message struct {
	Id        int64					`json:"id" bson:"_id"`
	Title     string    			`json:"title" bson:"title"`
	Body      string    			`json:"body" bson:"body"`
	CreatedAt time.Time 			`json:"created_at" bson:"createdAt"`
}

func (m *Message) Validate() error_utils.MessageErr {
	m.Title = strings.TrimSpace(m.Title)
	m.Body = strings.TrimSpace(m.Body)
	if m.Title == "" {
		return error_utils.NewUnprocessibleEntityError("Please enter a valid title")
	}
	if m.Body == "" {
		return error_utils.NewUnprocessibleEntityError("Please enter a valid body")
	}
	return nil
}