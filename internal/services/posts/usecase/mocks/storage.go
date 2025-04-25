package mocks

import "github.com/LeonidS635/soa/internal/services/posts/dto"

type PostsStorage struct {
	Post *dto.PostFullInfo
	Err  error
}

func NewPostsStorage() *PostsStorage {
	return &PostsStorage{}
}

func (s *PostsStorage) AddPost(UserId int32, post *dto.Post) (int32, error) {
	return s.Post.PostId, s.Err
}

func (s *PostsStorage) GetPost(postId, userId int32) (*dto.PostFullInfo, error) {
	return s.Post, s.Err
}

func (s *PostsStorage) GetAllPosts(page, postsPerPageN, userId int32, onlyAuthor bool) (
	[]*dto.PostFullInfo, error,
) {
	return []*dto.PostFullInfo{s.Post}, s.Err
}

func (s *PostsStorage) UpdatePost(postId, userId int32, newPost *dto.Post) error {
	return s.Err
}

func (s *PostsStorage) DeletePost(postId, userId int32) error {
	return s.Err
}
