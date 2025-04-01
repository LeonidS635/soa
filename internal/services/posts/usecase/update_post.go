package usecase

import "github.com/LeonidS635/soa/internal/services/posts/dto"

func (uc *PostsUseCase) UpdatePost(postId, userId int32, newPost *dto.Post) error {
	return uc.storage.UpdatePost(postId, userId, newPost)
}
