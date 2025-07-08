package mocks

import (
	"context"

	"github.com/LeonidS635/soa/internal/kafka/events"
)

type KafkaProducer struct{}

func (p *KafkaProducer) Publish(context.Context, string, events.Event) error {
	return nil
}
