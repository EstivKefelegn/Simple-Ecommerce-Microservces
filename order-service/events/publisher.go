package events

import (
	"encoding/json"
	"github/orderService/internal/domain"

	amqp "github.com/rabbitmq/amqp091-go"
)

type EventPublisher struct {
	Channel  *amqp.Channel
	Exchange string
}

func (p *EventPublisher) PublishOrderCreated(order *domain.Order) error {
	body, _ := json.Marshal(order)
	return p.Channel.Publish(
		p.Exchange,
		"order.created",
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		})
}
