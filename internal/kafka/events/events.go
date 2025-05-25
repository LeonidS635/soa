package events

import (
	"encoding/json"
	"time"
)

type EventType int

type Event struct {
	Moment time.Time `json:"moment"`
	Data   []byte    `json:"data"`
}

func createEvent(data any) (Event, error) {
	dataBytes, err := json.Marshal(data)
	if err != nil {
		return Event{}, err
	}

	return Event{
		Moment: time.Now(),
		Data:   dataBytes,
	}, nil
}

func UserRegistered(userID int32) (Event, error) {
	return createEvent(
		struct {
			UserID int32 `json:"user_id"`
		}{
			UserID: userID,
		},
	)
}

func PostViewed(userID, postID int32) (Event, error) {
	return createEvent(
		struct {
			UserID int32 `json:"user_id"`
			PostID int32 `json:"post_id"`
		}{
			UserID: userID,
			PostID: postID,
		},
	)
}
