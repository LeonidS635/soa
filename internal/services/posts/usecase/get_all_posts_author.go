package usecase

import "github.com/LeonidS635/soa/internal/services/posts/dto"

func (uc *PostsUseCase) GetAllPostsOfOneAuthor(authorId, postsPerPageN int32) ([][]*dto.PostFullInfo, error) {
	posts, err := uc.storage.GetAllPosts(postsPerPageN, authorId, true)
	if err != nil {
		return nil, err
	}
	return posts, nil
}
