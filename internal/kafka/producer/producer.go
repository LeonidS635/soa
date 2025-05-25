package producer

import (
	"context"
	"encoding/json"

	"github.com/LeonidS635/soa/internal/kafka/events"
	"github.com/google/uuid"
	"github.com/segmentio/kafka-go"
)

type KafkaProducer struct {
	kafka.Writer
}

func NewProducer(connString, topic string) *KafkaProducer {
	return &KafkaProducer{
		Writer: kafka.Writer{
			Addr:  kafka.TCP(connString),
			Topic: topic,
		},
	}
}

func (p *KafkaProducer) Publish(ctx context.Context, event events.Event) error {
	key, err := uuid.New().MarshalBinary()
	if err != nil {
		return err
	}

	value, err := json.Marshal(event)
	if err != nil {
		return err
	}

	return p.WriteMessages(
		ctx, kafka.Message{
			Key:   key,
			Value: value,
		},
	)
}

func (p *KafkaProducer) Close() error {
	return p.Writer.Close()
}
