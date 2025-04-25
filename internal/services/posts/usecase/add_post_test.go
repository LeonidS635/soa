package usecase

import (
	"errors"
	"testing"
	"time"

	"github.com/LeonidS635/soa/internal/services/posts/dto"
	"github.com/LeonidS635/soa/internal/services/posts/usecase/mocks"
	"github.com/stretchr/testify/require"
)

func TestPostsUseCase_AddPost(t *testing.T) {
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
		userId int32
		post   *dto.Post

		mock    mocks.PostsStorage
		wantErr bool
	}{
		{
			name:    "incorrect post",
			userId:  1,
			post:    nil,
			mock:    mocks.PostsStorage{},
			wantErr: true,
		},
		{
			name:   "storage error",
			userId: 1,
			post:   nil,
			mock: mocks.PostsStorage{
				Err: errors.New("something went wrong"),
			},
			wantErr: true,
		},
		{
			name:   "correct",
			userId: 1,
			post:   correctPost.Post,
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

				postId, err := uc.AddPost(tc.userId, tc.post)
				if (err != nil) != tc.wantErr {
					t.Errorf("AddPost() error = %v, wantErr %v", err, tc.wantErr)
					return
				}

				if err == nil {
					require.Equal(t, tc.userId, postId, "AddPost() returns wrong post id")
				}
			},
		)
	}
}
