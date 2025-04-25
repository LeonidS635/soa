package converters

import (
	"github.com/LeonidS635/soa/internal/pkg/services/postspb"
	"github.com/LeonidS635/soa/internal/services/posts/dto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func PostToPB(post *dto.Post) *postspb.Post {
	return &postspb.Post{
		IsPrivate: post.IsPrivate,
		Title:     post.Title,
		Tags:      post.Tags,
		Text:      post.Text,
	}
}

func PostServiceInfoToPB(details *dto.PostServiceInfo) *postspb.PostServiceInfo {
	return &postspb.PostServiceInfo{
		PostId:    details.PostId,
		AuthorId:  details.AuthorId,
		CreatedAt: timestamppb.New(details.CreatedAt),
		UpdatedAt: timestamppb.New(details.UpdatedAt),
	}
}

func PostFullInfoToPB(post *dto.PostFullInfo) *postspb.PostFullInfo {
	return &postspb.PostFullInfo{
		Post:    PostToPB(post.Post),
		Details: PostServiceInfoToPB(post.PostServiceInfo),
	}
}
