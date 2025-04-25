package usecase

import (
	"errors"
	"testing"
	"time"

	"github.com/LeonidS635/soa/internal/services/posts/dto"
	"github.com/LeonidS635/soa/internal/services/posts/usecase/mocks"
	"github.com/stretchr/testify/require"
)

func TestPostsUseCase_GetAllMyPosts(t *testing.T) {
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
		name          string
		page          int32
		postsPerPageN int32
		authorId      int32

		mock    mocks.PostsStorage
		wantErr bool
	}{
		{
			name:          "incorrect page",
			page:          -1,
			postsPerPageN: 10,
			authorId:      1,
			mock:          mocks.PostsStorage{},
			wantErr:       true,
		},
		{
			name:          "storage error",
			page:          1,
			postsPerPageN: 10,
			authorId:      1,
			mock: mocks.PostsStorage{
				Err: errors.New("something went wrong"),
			},
			wantErr: true,
		},
		{
			name:          "correct",
			page:          1,
			postsPerPageN: 10,
			authorId:      1,
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

				posts, err := uc.GetAllMyPosts(tc.page, tc.postsPerPageN, tc.authorId)
				if (err != nil) != tc.wantErr {
					t.Errorf("GetAllMyPosts() error = %v, wantErr %v", err, tc.wantErr)
					return
				}

				if err == nil {
					require.Equal(
						t, []*dto.PostFullInfo{&correctPost}, posts, "GetAllMyPosts() returns wrong posts list",
					)
				}
			},
		)
	}
}
