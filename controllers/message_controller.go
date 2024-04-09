package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"helloworld/models"
	"helloworld/services"
	error_utils "helloworld/utils"
	"net/http"
	"os"
	"strconv"
)

func CreateMessage(c *gin.Context) {
	var message models.Message
	if err := c.ShouldBindJSON(&message); err != nil {
		theErr := error_utils.NewUnprocessibleEntityError("invalid json body")
		c.JSON(theErr.Status(), theErr)
		return
	}
	err := services.CreateMessage(&message)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusCreated, message)
}

func getMessageId(msgIdParam string) (int64, error_utils.MessageErr) {
	msgId, err := strconv.ParseInt(msgIdParam, 10, 64)
	if err != nil {
		return 0, error_utils.NewBadRequestError("message id should be a number")
	}
	return msgId, nil
}


func UpdateMessage(c *gin.Context) {
	msgId, err := getMessageId(c.Param("message_id"))
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	var message models.Message
	if err := c.ShouldBindJSON(&message); err != nil {
		theErr := error_utils.NewUnprocessibleEntityError("invalid json body")
		c.JSON(theErr.Status(), theErr)
		return
	}
	message.Id = msgId
	err = services.UpdateMessage(&message)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, message)
}

func GetMessage(c *gin.Context) {
	res, err := services.GetMessage()
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, string(res))
}

func Hello(c *gin.Context) {
	newStr := NewString("Hello, World! ")
	c.String(200, newStr)
}

func HelloAppName(c *gin.Context) {
	host := os.Getenv("HOSTNAME")
	reply:= fmt.Sprintf("Hello from %s V7 \n", host)
	c.String(200, reply)
}

func Exit(c *gin.Context) {
	os.Exit(0)
}

func Health(c *gin.Context) {
	c.String(200, "Success")
}

func NewString(s string) string {
	return s+"12345"
}