package converters

import (
	"github.com/LeonidS635/soa/internal/pkg/services/postspb"
	"github.com/LeonidS635/soa/internal/services/posts/dto"
)

func PostToDTO(post *postspb.Post) *dto.Post {
	return &dto.Post{
		IsPrivate: post.GetIsPrivate(),
		Title:     post.GetTitle(),
		Tags:      post.GetTags(),
		Text:      post.GetText(),
	}
}

func PostServiceInfoToDTO(details *postspb.PostServiceInfo) *dto.PostServiceInfo {
	return &dto.PostServiceInfo{
		PostId:    details.GetPostId(),
		AuthorId:  details.GetAuthorId(),
		CreatedAt: details.GetCreatedAt().AsTime(),
		UpdatedAt: details.GetUpdatedAt().AsTime(),
	}
}

func PostFullInfoToDTO(post *postspb.PostFullInfo) *dto.PostFullInfo {
	return &dto.PostFullInfo{
		Post:            PostToDTO(post.GetPost()),
		PostServiceInfo: PostServiceInfoToDTO(post.GetDetails()),
	}
}
