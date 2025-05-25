package dto

import "time"

type PostStats struct {
	Views    int32
	Likes    int32
	Comments int32
}

type PostDynamics struct {
	Stats []struct {
		Day time.Time
		PostStats
	}
}
