package usecase

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/gabrielmellooliveira/m2c-digital-consumer-golang/internal/domain/models"
	"github.com/gabrielmellooliveira/m2c-digital-consumer-golang/internal/infra/database"
	"github.com/gabrielmellooliveira/m2c-digital-consumer-golang/internal/infra/http"
	"github.com/gabrielmellooliveira/m2c-digital-consumer-golang/internal/infra/queue"
)

type ConsumeMessagesUseCase struct {
	MongoDbAdapter  database.MongoDbAdapter
	HttpAdapter     http.HttpAdapter
	RabbitMqAdapter queue.RabbitMqAdapter
	RedisAdapter    database.RedisAdapter
}

func NewConsumeMessagesUseCase(
	mongoDbAdapter database.MongoDbAdapter,
	httpAdapter http.HttpAdapter,
	rabbitmqAdapter queue.RabbitMqAdapter,
	redisAdapter database.RedisAdapter,
) *ConsumeMessagesUseCase {
	return &ConsumeMessagesUseCase{
		MongoDbAdapter:  mongoDbAdapter,
		HttpAdapter:     httpAdapter,
		RabbitMqAdapter: rabbitmqAdapter,
		RedisAdapter:    redisAdapter,
	}
}

func (u *ConsumeMessagesUseCase) Execute() {
	const M2C_DIGITAL_QUEUE_NAME = "m2c_digital_messages_queue"

	u.RabbitMqAdapter.Consume(M2C_DIGITAL_QUEUE_NAME, func(data []byte) error {
		fmt.Printf("Message: %s\n", data)

		message, err := models.CreateMessage(data)
		if err != nil {
			return err
		}

		err = u.MongoDbAdapter.Insert("messages", message.Prepare())
		if err != nil {
			return err
		}

		campaignKey := u.getCampaignKey(message.CampaignId)

		wereRead, err := u.haveAllMessagesBeenRead(campaignKey, message.Total)
		if err != nil {
			return err
		}

		if wereRead {
			err := u.updateCampaignStatusToSent(message.CampaignId)
			if err != nil {
				return err
			}

			err = u.RedisAdapter.Delete(campaignKey)
			if err != nil {
				return err
			}
		}

		return nil
	})
}

func (u *ConsumeMessagesUseCase) getCampaignKey(campaignId string) string {
	return "campaign:" + campaignId + ":count"
}

func (u *ConsumeMessagesUseCase) haveAllMessagesBeenRead(campaignKey string, total int) (bool, error) {
	u.RedisAdapter.Increment(campaignKey)
	messageCount, err := u.RedisAdapter.Get(campaignKey)
	if err != nil {
		return false, errors.New("Failed to get value in Redis: " + err.Error())
	}

	parsedCount, err := strconv.Atoi(messageCount)
	if err != nil {
		return false, errors.New("Failed to convert string to integer: " + err.Error())
	}

	return total == parsedCount, nil
}

func (u *ConsumeMessagesUseCase) updateCampaignStatusToSent(campaignId string) error {
	_, err := u.HttpAdapter.Put("/campaigns/sent/"+campaignId, nil)
	return err
}
