package kafka

import (
	"fmt"
	"testing"
	"time"

	"github.com/LeonidS635/soa/internal/kafka"
	"github.com/LeonidS635/soa/internal/kafka/events"
	"github.com/stretchr/testify/require"
)

func TestStatisticsService_Kafka(t *testing.T) {
	p := kafka.NewProducer(kafkaAddr)
	defer func() {
		if err := p.CloseAll(); err != nil {
			t.Fatal(err)
		}
	}()

	p.RegisterTopic(postsViewsTopic)
	p.RegisterTopic(postsLikesTopic)
	p.RegisterTopic(postsCommentsTopic)

	publish := func(topic string, postID, userID int32) error {
		return p.Publish(
			t.Context(), topic, &events.PostViewedEvent{
				Moment: time.Now(),
				PostID: postID,
				UsedID: userID,
			},
		)
	}

	const postsN, viewsN, likesN, commentsN = 3, 2, 2, 2

	for post := int32(0); post < postsN; post++ {
		for i := 0; i < viewsN; i++ {
			require.NoError(t, publish(postsViewsTopic, post, post))
		}
		for i := 0; i < likesN; i++ {
			require.NoError(t, publish(postsLikesTopic, post, post))
		}
		for i := 0; i < commentsN; i++ {
			require.NoError(t, publish(postsCommentsTopic, post, post))
		}
	}

	require.Error(t, publish("some-unregistered-topic", 123, 123))

	p.RegisterTopic("some-unknown-topic")
	require.Error(t, publish("some-unknown-topic", 456, 456))

	tests := []struct {
		name     string
		metric   string
		expected uint64
	}{
		{
			name:     "views",
			metric:   "views",
			expected: viewsN,
		},
		{
			name:     "likes",
			metric:   "likes",
			expected: likesN,
		},
		{
			name:     "comments",
			metric:   "comments",
			expected: commentsN,
		},
	}
	for _, test := range tests {
		t.Run(
			test.name, func(t *testing.T) {
				for post := int32(0); post < postsN; post++ {
					metricValue, err := getMetricValue(t.Context(), post, test.metric)
					require.NoError(t, err, fmt.Sprintf("post %d", post))
					require.Equal(t, test.expected, metricValue, fmt.Sprintf("post %d", post))
				}
			},
		)
	}
}
