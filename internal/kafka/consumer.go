package kafka

import (
	"context"
	"fmt"

	"github.com/segmentio/kafka-go"
)

type Consumer struct {
	connString string
	consumers  map[string]*kafka.Reader
}

func NewConsumer(connString string) *Consumer {
	return &Consumer{
		connString: connString,
		consumers:  make(map[string]*kafka.Reader),
	}
}

func (c *Consumer) RegisterTopic(topic string) {
	c.consumers[topic] = kafka.NewReader(
		kafka.ReaderConfig{
			Brokers: []string{c.connString},
			Topic:   topic,
		},
	)
}

func (c *Consumer) Consume(ctx context.Context, topic string, handler func(context.Context, []byte) error) error {
	r, ok := c.consumers[topic]
	if !ok {
		return fmt.Errorf("error consuming event: unknown topic %q", topic)
	}

	for {
		m, err := r.ReadMessage(ctx)
		if err != nil {
			return fmt.Errorf("error consuming event: error reading message: %w", err)
		}

		if err = handler(ctx, m.Value); err != nil {
			return err
		}
	}
}

func (c *Consumer) Close(topic string) error {
	if consumer, ok := c.consumers[topic]; ok {
		return consumer.Close()
	}
	return fmt.Errorf("error closing consumer: unknown topic %q", topic)
}

func (c *Consumer) CloseAll() error {
	var err error
	for _, consumer := range c.consumers {
		if closeErr := consumer.Close(); closeErr != nil {
			err = closeErr
		}
	}
	return err
}
