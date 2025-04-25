package validators

import (
	"errors"

	"github.com/LeonidS635/soa/internal/services/posts/dto"
)

func Post(post *dto.Post) error {
	if post == nil {
		return errors.New("post is empty")
	}

	if len(post.Title) == 0 {
		return errors.New("post title is empty")
	}
	if len(post.Title) > 100 {
		return errors.New("post title is too long")
	}

	if len(post.Tags) > 10 {
		return errors.New("post has too many tags")
	}
	for _, tag := range post.Tags {
		if len(tag) == 0 {
			return errors.New("tag is empty")
		}
		if len(tag) > 100 {
			return errors.New("tag is too long")
		}
	}

	if len(post.Text) == 0 {
		return errors.New("post text is empty")
	}
	if len(post.Text) > 5000 {
		return errors.New("post text is too long")
	}

	return nil
}
