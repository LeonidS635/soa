package usecase

import "github.com/LeonidS635/soa/internal/services/posts/dto"

func (uc *PostsUseCase) GetPost(postId, userId int32) (*dto.PostFullInfo, error) {
	return uc.storage.GetPost(postId, userId)
}
