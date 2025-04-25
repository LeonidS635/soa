package usecase

import (
	"github.com/LeonidS635/soa/internal/services/posts/dto"
	"github.com/LeonidS635/soa/internal/services/posts/usecase/validators"
)

func (uc *PostsUseCase) GetAllMyPosts(page, postsPerPageN, authorId int32) ([]*dto.PostFullInfo, error) {
	if err := validators.Page(page); err != nil {
		return nil, err
	}
	return uc.storage.GetAllPosts(page, postsPerPageN, authorId, true)
}
