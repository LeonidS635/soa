package events

import (
	"time"
)

type EventType int

type Event struct {
	Moment time.Time `json:"moment"`
	Data   any       `json:"data"`
}

func createEvent(data any) Event {
	return Event{
		Moment: time.Now(),
		Data:   data,
	}
}

func UserRegistered(userID int32) Event {
	return createEvent(
		struct {
			UserID int32 `json:"user_id"`
		}{
			UserID: userID,
		},
	)
}

func PostViewed(userID, postID int32) Event {
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

func PostLiked(userID, postID int32) Event {
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

func PostCommented(userID, postID int32) Event {
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
