package events

import (
	"encoding/json"
	"time"
)

type PostCommentedEvent struct {
	Moment    time.Time `json:"moment"`
	PostID    int32     `json:"post_id"`
	CommentID int32     `json:"comment_id"`
	UsedID    int32     `json:"used_id"`
}

func (e *PostCommentedEvent) Marshal() ([]byte, error) {
	return json.Marshal(*e)
}

func (e *PostCommentedEvent) Unmarshal(b []byte) error {
	return json.Unmarshal(b, e)
}
