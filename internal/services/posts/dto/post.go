package dto

import "time"

type Post struct {
	IsPrivate bool `json:"is_private"`

	Title string   `json:"title"`
	Tags  []string `json:"tags"`
	Text  string   `json:"text"`
}

type PostServiceInfo struct {
	PostId   int32 `json:"post_id"`
	AuthorId int32 `json:"author_id"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type PostFullInfo struct {
	*Post
	*PostServiceInfo
}
