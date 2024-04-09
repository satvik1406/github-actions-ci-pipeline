package services

import (
	"helloworld/domain"
	"helloworld/models"
	error_utils "helloworld/utils"
	"time"
)

func CreateMessage(message *models.Message)  error_utils.MessageErr {
	if err := message.Validate(); err != nil {
		return err
	}
	message.CreatedAt = time.Now()
	err := domain.Create(message)
	if err != nil {
		return err
	}
	return nil
}

func UpdateMessage(message *models.Message) error_utils.MessageErr {
	if err := message.Validate(); err != nil {
		return err
	}
	err := domain.Update(message)
	if err != nil {
		return err
	}
	return nil
}

func GetMessage() ([]byte, error_utils.MessageErr) {
	res, err := domain.Get()
	if err != nil {
		return nil, err
	}
	return res,nil
}