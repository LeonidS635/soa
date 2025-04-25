package usecase

import (
	"github.com/LeonidS635/soa/internal/services/posts/dto"
	"github.com/LeonidS635/soa/internal/services/posts/usecase/validators"
)

func (uc *PostsUseCase) UpdatePost(postId, userId int32, newPost *dto.Post) error {
	if err := validators.Post(newPost); err != nil {
		return err
	}
	return uc.storage.UpdatePost(postId, userId, newPost)
}
