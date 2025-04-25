package usecase

import (
	"errors"
	"testing"
	"time"

	"github.com/LeonidS635/soa/internal/services/posts/dto"
	"github.com/LeonidS635/soa/internal/services/posts/usecase/mocks"
	"github.com/stretchr/testify/require"
)

func TestPostsUseCase_GetPost(t *testing.T) {
	mockStorage := mocks.NewPostsStorage()
	uc := NewPostsUseCase(mockStorage)

	correctPost := dto.PostFullInfo{
		Post: &dto.Post{
			IsPrivate: false,
			Title:     "some title",
			Tags:      []string{"tag1", "tag2"},
			Text:      "some text",
		},
		PostServiceInfo: &dto.PostServiceInfo{
			PostId:    1,
			AuthorId:  1,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	testCases := []struct {
		name   string
		postId int32
		userId int32

		mock    mocks.PostsStorage
		wantErr bool
	}{
		{
			name:   "storage error",
			postId: 1,
			userId: 1,
			mock: mocks.PostsStorage{
				Err: errors.New("something went wrong"),
			},
			wantErr: true,
		},
		{
			name:   "correct",
			postId: 1,
			userId: 1,
			mock: mocks.PostsStorage{
				Post: &correctPost,
			},
			wantErr: false,
		},
	}

	for _, tc := range testCases {
		t.Run(
			tc.name, func(t *testing.T) {
				mockStorage.Post = tc.mock.Post
				mockStorage.Err = tc.mock.Err

				post, err := uc.GetPost(tc.postId, tc.userId)
				if (err != nil) != tc.wantErr {
					t.Errorf("GetPost() error = %v, wantErr %v", err, tc.wantErr)
					return
				}

				if err == nil {
					require.Equal(t, &correctPost, post, "GetPost() returns wrong post")
				}
			},
		)
	}
}
