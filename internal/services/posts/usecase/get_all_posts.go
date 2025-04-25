package usecase

import (
	"github.com/LeonidS635/soa/internal/services/posts/dto"
	"github.com/LeonidS635/soa/internal/services/posts/usecase/validators"
)

func (uc *PostsUseCase) GetAllPosts(page, postsPerPageN int32) ([]*dto.PostFullInfo, error) {
	if err := validators.Page(page); err != nil {
		return nil, err
	}
	return uc.storage.GetAllPosts(page, postsPerPageN, -1, false)
}
