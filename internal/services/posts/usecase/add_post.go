package usecase

import (
	"github.com/LeonidS635/soa/internal/services/posts/dto"
	"github.com/LeonidS635/soa/internal/services/posts/usecase/validators"
)

func (uc *PostsUseCase) AddPost(userId int32, post *dto.Post) (int32, error) {
	if err := validators.Post(post); err != nil {
		return -1, err
	}

	postId, err := uc.storage.AddPost(userId, post)
	if err != nil {
		return -1, err
	}
	return postId, nil
}
