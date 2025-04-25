package validators

import (
	"strings"
	"testing"

	"github.com/LeonidS635/soa/internal/services/posts/dto"
)

func TestPost(t *testing.T) {
	testCases := []struct {
		name    string
		post    dto.Post
		wantErr bool
	}{
		{
			name:    "empty",
			post:    dto.Post{},
			wantErr: true,
		},
		{
			name: "empty title",
			post: dto.Post{
				Text: "some text",
			},
			wantErr: true,
		},
		{
			name: "long title",
			post: dto.Post{
				Title: strings.Repeat("long title", 100),
				Text:  "some text",
			},
			wantErr: true,
		},
		{
			name: "many tags",
			post: dto.Post{
				Title: "title",
				Tags:  make([]string, 1000),
				Text:  "some text",
			},
			wantErr: true,
		},
		{
			name: "empty tag",
			post: dto.Post{
				Title: "title",
				Tags:  []string{""},
				Text:  "some text",
			},
			wantErr: true,
		},
		{
			name: "long tag",
			post: dto.Post{
				Title: "title",
				Tags:  []string{strings.Repeat("long tag", 100)},
				Text:  "some text",
			},
			wantErr: true,
		},
		{
			name: "empty text",
			post: dto.Post{
				Title: "title",
				Tags:  []string{"some", "tags"},
			},
			wantErr: true,
		},
		{
			name: "long text",
			post: dto.Post{
				Title: "title",
				Tags:  []string{"some", "tags"},
				Text:  strings.Repeat("long text", 1000),
			},
			wantErr: true,
		},
		{
			name: "correct",
			post: dto.Post{
				Title: "title",
				Tags:  []string{"some", "tags"},
				Text:  "some text",
			},
			wantErr: false,
		},
	}

	for _, tc := range testCases {
		t.Run(
			tc.name, func(t *testing.T) {
				err := Post(&tc.post)
				if (err != nil) != tc.wantErr {
					t.Errorf("Post() error = %v, wantErr %v", err, tc.wantErr)
				}
			},
		)
	}
}
