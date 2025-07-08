package events

import (
	"encoding/json"
	"time"
)

type PostLikedEvent struct {
	Moment time.Time `json:"moment"`
	PostID int32     `json:"post_id"`
	UsedID int32     `json:"used_id"`
}

func (e *PostLikedEvent) Marshal() ([]byte, error) {
	return json.Marshal(*e)
}

func (e *PostLikedEvent) Unmarshal(b []byte) error {
	return json.Unmarshal(b, e)
}
