package integration

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"helloworld/controllers"
	"log"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)
func TestCreateMessage(t *testing.T) {

	database()

	gin.SetMode(gin.TestMode)

	err := refreshMessagesTable()
	if err != nil {
		log.Fatal(err)
	}
	samples := []struct {
		inputJSON  string
		statusCode int
		title      string
		body       string
		errMessage string
	}{
		{
			inputJSON:  `{"id": 0, "title":"the title", "body": "the body"}`,
			statusCode: 201,
			title:      "the title",
			body:       "the body",
			errMessage: "",
		},
		{
			inputJSON:  `{"id": 1, "title":"the title", "body": "the body"}`,
			statusCode: 201,
			title:      "the title",
			body:       "the body",
			errMessage: "",
		},
		{
			inputJSON:  `{"id": 2, "title":"", "body": "the body"}`,
			statusCode: 422,
			errMessage: "Please enter a valid title",
		},
		{
			inputJSON:  `{"id": 3, "title":"the title", "body": ""}`,
			statusCode: 422,
			errMessage: "Please enter a valid body",
		},
		{
			//when an integer is used like a string for title
			inputJSON:  `{"id": 4, "title": 12345, "body": "the body"}`,
			statusCode: 422,
			errMessage: "invalid json body",
		},
		{
			//when an integer is used like a string for body
			inputJSON:  `{"id": 5, "title": "the title", "body": 123453 }`,
			statusCode: 422,
			errMessage: "invalid json body",
		},
	}
	for _, v := range samples {
		r := gin.Default()
		r.POST("/messages", controllers.CreateMessage)
		req, err := http.NewRequest(http.MethodPost, "/messages", bytes.NewBufferString(v.inputJSON))
		if err != nil {
			t.Errorf("this is the error: %v\n", err)
		}
		rr := httptest.NewRecorder()
		r.ServeHTTP(rr, req)

		responseMap := make(map[string]interface{})
		err = json.Unmarshal(rr.Body.Bytes(), &responseMap)
		if err != nil {
			t.Errorf("Cannot convert to json: %v", err)
		}
		fmt.Println("this is the response data: ", responseMap)
		assert.Equal(t, rr.Code, v.statusCode)
		if v.statusCode == 201 {
			//casting the interface to map:
			assert.Equal(t, responseMap["title"], v.title)
			assert.Equal(t, responseMap["body"], v.body)
		}
		if v.statusCode == 400 || v.statusCode == 422 || v.statusCode == 500 && v.errMessage != "" {
			assert.Equal(t, responseMap["message"], v.errMessage)
		}
	}
}

func TestUpdateMessage(t *testing.T) {

	database()

	gin.SetMode(gin.TestMode)

	err := refreshMessagesTable()
	if err != nil {
		log.Fatal(err)
	}
	message, err := seedOneMessage()
	if err != nil {
		t.Errorf("Error while seeding table: %s", err)
	}

	//Get only the first message id
	firstId := message.Id

	samples := []struct {
		id          string
		inputJSON  string
		statusCode int
		title      string
		body       string
		errMessage string
	}{
		{
			id:          strconv.Itoa(int(firstId)),
			inputJSON:  `{"title":"update title", "body": "update body"}`,
			statusCode: 200,
			title:      "update title",
			body:       "update body",
			errMessage: "",
		},
		{
			// "second title" belongs to the second message so, the cannot be used for the first message
			id:         strconv.Itoa(int(firstId)),
			inputJSON:  `{"title":"second title", "body": "update body"}`,
			statusCode: 200,
			title:      "second title",
			body:       "update body",
			errMessage: "",
		},
		{
			//Empty title
			id:          strconv.Itoa(int(firstId)),
			inputJSON:  `{"title":"", "body": "update body"}`,
			statusCode: 422,
			errMessage: "Please enter a valid title",
		},
		{
			//Empty body
			id:          strconv.Itoa(int(firstId)),
			inputJSON:  `{"title":"the title", "body": ""}`,
			statusCode: 422,
			errMessage: "Please enter a valid body",
		},
		{
			//when an integer is used like a string for title
			id:          strconv.Itoa(int(firstId)),
			inputJSON:  `{"title": 12345, "body": "the body"}`,
			statusCode: 422,
			errMessage: "invalid json body",
		},
		{
			//when an integer is used like a string for body
			id:          strconv.Itoa(int(firstId)),
			inputJSON:  `{"title": "the title", "body": 123453 }`,
			statusCode: 422,
			errMessage: "invalid json body",
		},
		{
			id:         "unknown",
			statusCode: 400,
			errMessage: "message id should be a number",
		},
		{
			id:         strconv.Itoa(12322), //an id that does not exist
			inputJSON:  `{"title":"the title", "body": "the body"}`,
			statusCode: 404,
			errMessage: "no record matching given id",
		},
	}
	for _, v := range samples {
		r := gin.Default()
		r.PUT("/messages/:message_id", controllers.UpdateMessage)
		req, err := http.NewRequest(http.MethodPut, "/messages/"+v.id, bytes.NewBufferString(v.inputJSON))
		if err != nil {
			t.Errorf("this is the error: %v\n", err)
		}
		rr := httptest.NewRecorder()
		r.ServeHTTP(rr, req)

		responseMap := make(map[string]interface{})
		err = json.Unmarshal(rr.Body.Bytes(), &responseMap)
		if err != nil {
			t.Errorf("Cannot convert to json: %v", err)
		}
		assert.Equal(t, rr.Code, v.statusCode)
		if v.statusCode == 200 {
			//casting the interface to map:
			assert.Equal(t, responseMap["title"], v.title)
			assert.Equal(t, responseMap["body"], v.body)
		}
		if v.statusCode == 400 || v.statusCode == 422 || v.statusCode == 500 && v.errMessage != "" {
			assert.Equal(t, responseMap["message"], v.errMessage)
		}
	}
}