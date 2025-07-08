package e2e

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestE2E_Post(t *testing.T) {
	t.Cleanup(
		func() {
			if err := dbm.clearDB(context.Background()); err != nil {
				t.Error(err)
			}
		},
	)

	cookieJar, err := cookiejar.New(nil)
	if err != nil {
		t.Fatal(err)
	}

	client := http.Client{
		Transport: http.DefaultTransport,
		Jar:       cookieJar,
		Timeout:   time.Second * 10,
	}

	// Registering a new user

	registerBody := struct {
		Username string `json:"username"`
		Password string `json:"password"`
		Email    string `json:"email"`
	}{
		Username: "John Doe",
		Password: "secure_password123",
		Email:    "johndoe@example.com",
	}

	body, err := json.Marshal(&registerBody)
	if err != nil {
		t.Fatal(err)
	}
	registerResp, err := client.Post(baseURL+"/signup", "application/json", bytes.NewReader(body))
	if err != nil {
		t.Fatal(err)
	}
	defer registerResp.Body.Close()

	require.Equal(t, http.StatusCreated, registerResp.StatusCode)

	// Making a new post

	postBody := struct {
		IsPrivate bool     `json:"is_private"`
		Title     string   `json:"title"`
		Tags      []string `json:"tags"`
		Text      string   `json:"text"`
	}{
		IsPrivate: false,
		Title:     "title",
		Tags:      []string{"tag1", "tag2"},
		Text:      "text",
	}
	body, err = json.Marshal(&postBody)
	if err != nil {
		t.Fatal(err)
	}
	postResp, err := client.Post(baseURL+"/posts", "application/json", bytes.NewReader(body))
	if err != nil {
		t.Fatal(err)
	}
	defer postResp.Body.Close()

	require.Equal(t, http.StatusOK, postResp.StatusCode)

	var postID struct {
		ID int `json:"postId"`
	}
	require.NoError(t, json.NewDecoder(postResp.Body).Decode(&postID))

	// Getting created post

	getPostResp, err := client.Get(baseURL + fmt.Sprintf("/posts/%d", postID.ID))
	if err != nil {
		t.Fatal(err)
	}
	defer getPostResp.Body.Close()

	require.Equal(t, http.StatusOK, getPostResp.StatusCode)

	var post struct {
		PostFullInfo struct {
			Post struct {
				IsPrivate bool     `json:"isPrivate"`
				Title     string   `json:"title"`
				Tags      []string `json:"tags"`
				Text      string   `json:"text"`
			} `json:"post"`

			Details struct {
				PostID    int       `json:"postId"`
				UserID    int       `json:"userId"`
				CreatedAt time.Time `json:"createdAt"`
				UpdatedAt time.Time `json:"updatedAt"`
			} `json:"details"`
		} `json:"post"`
	}
	require.NoError(t, json.NewDecoder(getPostResp.Body).Decode(&post))

	require.Equal(t, postID.ID, post.PostFullInfo.Details.PostID)
	require.Equal(t, postBody.IsPrivate, post.PostFullInfo.Post.IsPrivate)
	require.Equal(t, postBody.Title, post.PostFullInfo.Post.Title)
	require.Equal(t, postBody.Tags, post.PostFullInfo.Post.Tags)
	require.Equal(t, postBody.Text, post.PostFullInfo.Post.Text)
}
