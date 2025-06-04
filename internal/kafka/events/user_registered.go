package events

import (
	"encoding/json"
	"time"
)

type UserRegisteredEvent struct {
	Moment time.Time `json:"moment"`
	UsedID int32     `json:"used_id"`
}

func (e *UserRegisteredEvent) Marshal() ([]byte, error) {
	return json.Marshal(*e)
}

func (e *UserRegisteredEvent) Unmarshal(b []byte) error {
	return json.Unmarshal(b, e)
}
