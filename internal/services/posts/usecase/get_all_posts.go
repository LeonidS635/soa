package usecase

import "github.com/LeonidS635/soa/internal/services/posts/dto"

func (uc *PostsUseCase) GetAllPosts(postsPerPageN int32) ([][]*dto.PostFullInfo, error) {
	posts, err := uc.storage.GetAllPosts(postsPerPageN, -1, false)
	if err != nil {
		return nil, err
	}
	return posts, nil
}
