package handlers

import (
	"github.com/LeonidS635/soa/internal/pkg/services/postspb"
	"github.com/LeonidS635/soa/internal/services/posts/usecase"
)

type PostsHandlers struct {
	useCase *usecase.PostsUseCase
	postspb.UnimplementedPostsServiceServer
}

func NewPostsHandlers(useCase *usecase.PostsUseCase) *PostsHandlers {
	return &PostsHandlers{useCase: useCase}
}
