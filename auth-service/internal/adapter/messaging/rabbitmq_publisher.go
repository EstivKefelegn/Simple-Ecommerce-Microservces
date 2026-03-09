package messaging

import (
	"context"
	"encoding/json"
	"github/ecommerceMSAuth/internal/domain"
	"github/ecommerceMSAuth/pkg"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

const (
	UserExchange          = "user_events"
	UserCreatedRoutingKey = "user.created"
)

type RabbitPublisher struct {
	channel  *amqp.Channel
	exchange string
}

func NewRabbitMQPublisher(conn *amqp.Connection) (*RabbitPublisher, error) {
	ch, err := conn.Channel()
	if err != nil {
		pkg.ErrorHandler(err, "Message broker channel error")
	}

	err = ch.ExchangeDeclare(
		UserExchange,
		"topic",
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		pkg.ErrorHandler(err, "message broker error")
	}
	return &RabbitPublisher{
		channel:  ch,
		exchange: UserExchange,
	}, nil
}

func (p *RabbitPublisher) PublishUserCreated(ctx context.Context, event domain.UserCreatedEvent) error {
	body, err := json.Marshal(event)
	if err != nil {
		return err
	}

	err = p.channel.PublishWithContext(ctx, p.exchange, UserCreatedRoutingKey, false, false,
	 amqp.Publishing{
		ContentType: "application/json",
		DeliveryMode: amqp.Persistent,
		Type: UserCreatedRoutingKey,
		Body: body,
	 })

	if err != nil {
		return err
	} 

	log.Printf("event published: %s user=%s", UserCreatedRoutingKey, event.ID)

	return nil
}

func (p *RabbitPublisher) Close() error {
	return p.channel.Close()
}
