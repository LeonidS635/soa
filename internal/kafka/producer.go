package kafka

import (
	"context"
	"fmt"

	"github.com/LeonidS635/soa/internal/kafka/events"
	"github.com/google/uuid"
	"github.com/segmentio/kafka-go"
)

type Producer struct {
	connString string
	producers  map[string]*kafka.Writer
}

func NewProducer(connString string) *Producer {
	return &Producer{
		connString: connString,
		producers:  make(map[string]*kafka.Writer),
	}
}

func (p *Producer) RegisterTopic(topic string) {
	p.producers[topic] = &kafka.Writer{
		Addr:  kafka.TCP(p.connString),
		Topic: topic,
	}
}

func (p *Producer) Publish(ctx context.Context, topic string, event events.Event) error {
	if _, ok := p.producers[topic]; !ok {
		return fmt.Errorf("error publishing event: unknown topic %q", topic)
	}

	key, err := uuid.New().MarshalBinary()
	if err != nil {
		return err
	}

	value, err := event.Marshal()
	if err != nil {
		return err
	}

	return p.producers[topic].WriteMessages(
		ctx, kafka.Message{
			Key:   key,
			Value: value,
		},
	)
}

func (p *Producer) Close(topic string) error {
	if producer, ok := p.producers[topic]; ok {
		return producer.Close()
	}
	return fmt.Errorf("error closing producer: unknown topic %q", topic)
}

func (p *Producer) CloseAll() error {
	var err error
	for _, producer := range p.producers {
		if closeErr := producer.Close(); closeErr != nil {
			err = closeErr
		}
	}
	return err
}
