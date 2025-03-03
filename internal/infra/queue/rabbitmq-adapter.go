package queue

import (
	"log"

	"github.com/rabbitmq/amqp091-go"
)

type ComplexMessage struct {
	Identifier  string `json:"identifier"`
	Message     string `json:"message"`
	PhoneNumber string `json:"phoneNumber"`
	CampaignId  string `json:"campaignId"`
}

type RabbitMqAdapter struct {
	Url string
}

func NewRabbitMqAdapter(url string) *RabbitMqAdapter {
	return &RabbitMqAdapter{
		Url: url,
	}
}

type Handler func(body []byte) error

func (e *RabbitMqAdapter) Consume(queueName string, handler Handler) {
	connection := e.getConnection()
	defer connection.Close()

	channel := e.getChannel(connection)
	defer channel.Close()

	delivery := e.getDelivery(channel, queueName)

	done := make(chan bool)

	go func() {
		for message := range delivery {
			err := handler(message.Body)
			if err != nil {
				println(err.Error())
			}
		}
	}()

	<-done
}

func (e *RabbitMqAdapter) getConnection() *amqp091.Connection {
	conn, err := amqp091.Dial(e.Url)
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}

	return conn
}

func (e *RabbitMqAdapter) getChannel(connection *amqp091.Connection) *amqp091.Channel {
	ch, err := connection.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %v", err)
	}

	return ch
}

func (e *RabbitMqAdapter) getDelivery(channel *amqp091.Channel, queueName string) <-chan amqp091.Delivery {
	msgs, err := channel.Consume(
		queueName,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Failed to register a consumer: %v", err)
	}

	return msgs
}
