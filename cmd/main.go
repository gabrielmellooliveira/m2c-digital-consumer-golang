package main

import (
	"github.com/gabrielmellooliveira/m2c-digital-consumer-golang/configs"
	"github.com/gabrielmellooliveira/m2c-digital-consumer-golang/internal/infra/database"
	"github.com/gabrielmellooliveira/m2c-digital-consumer-golang/internal/infra/http"
	"github.com/gabrielmellooliveira/m2c-digital-consumer-golang/internal/infra/queue"
	"github.com/gabrielmellooliveira/m2c-digital-consumer-golang/internal/usecase"
)

const M2C_DIGITAL_DATABASE_NAME = "m2c_digital_db"

func main() {
	config, err := configs.LoadConfig()
	if err != nil {
		panic(err)
	}

	mongoDbAdapter := database.NewMongoDbAdapter(config.MongoDBUrl, M2C_DIGITAL_DATABASE_NAME)
	httpAdapter := http.NewHttpAdapter(config.M2CDigitalApiUrl)
	rabbitMqAdapter := queue.NewRabbitMqAdapter(config.RabbitMQUrl)
	redisAdapter := database.NewRedisAdapter(config.RedisUrl)

	mongoDbAdapter.Connect()
	redisAdapter.Connect()

	httpAdapter.AddHeader("x-api-key", config.M2CDigitalApiKey)

	consumeMessageUseCase := usecase.NewConsumeMessagesUseCase(*mongoDbAdapter, *httpAdapter, *rabbitMqAdapter, *redisAdapter)
	consumeMessageUseCase.Execute()
}
