package usecase

import (
	"errors"
	"testing"

	"github.com/LeonidS635/soa/internal/services/posts/usecase/mocks"
)

func TestPostsUseCase_DeletePost(t *testing.T) {
	mockStorage := mocks.NewPostsStorage()
	uc := NewPostsUseCase(mockStorage)

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
			name:    "correct",
			postId:  1,
			userId:  1,
			mock:    mocks.PostsStorage{},
			wantErr: false,
		},
	}

	for _, tc := range testCases {
		t.Run(
			tc.name, func(t *testing.T) {
				mockStorage.Err = tc.mock.Err

				err := uc.DeletePost(tc.postId, tc.userId)
				if (err != nil) != tc.wantErr {
					t.Errorf("DeletePost() error = %v, wantErr %v", err, tc.wantErr)
				}
			},
		)
	}
}
