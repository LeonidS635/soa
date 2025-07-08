package e2e

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/cookiejar"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestE2E_Stats(t *testing.T) {
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

	const viewsN, likesN, commentsN = 5, 4, 2

	// Get post

	for i := 0; i < viewsN; i++ {
		getPostResp, err := client.Get(baseURL + fmt.Sprintf("/posts/%d", postID.ID))
		if err != nil {
			t.Fatal(err)
		}
		_ = getPostResp.Body.Close()

		require.Equal(t, http.StatusOK, getPostResp.StatusCode)
	}

	// Like post

	for i := 0; i < likesN; i++ {
		likeResp, err := client.Post(baseURL+fmt.Sprintf("/posts/%d/like", postID.ID), "", nil)
		if err != nil {
			t.Fatal(err)
		}
		_ = likeResp.Body.Close()

		require.Equal(t, http.StatusOK, likeResp.StatusCode)
	}

	// Comment post

	var comment struct {
		Text string `json:"text"`
	}
	for i := 0; i < commentsN; i++ {
		comment.Text = fmt.Sprintf("comment %d", i)

		body, _ = json.Marshal(&comment)
		commentResp, err := client.Post(
			baseURL+fmt.Sprintf("/posts/%d/comment", postID.ID), "application/json", bytes.NewReader(body),
		)
		if err != nil {
			t.Fatal(err)
		}
		_ = commentResp.Body.Close()

		require.Equal(t, http.StatusOK, commentResp.StatusCode)
	}

	// Get statistics

	statsResp, err := client.Get(baseURL + fmt.Sprintf("/posts/%d/stats", postID.ID))
	if err != nil {
		t.Fatal(err)
	}
	defer statsResp.Body.Close()

	b, _ := io.ReadAll(statsResp.Body)
	t.Log(string(b))

	require.Equal(t, http.StatusOK, statsResp.StatusCode)

	var stats struct {
		Views    int `json:"views"`
		Likes    int `json:"likes"`
		Comments int `json:"comments"`
	}
	require.NoError(t, json.NewDecoder(statsResp.Body).Decode(&stats))

	require.Equal(t, viewsN, stats.Views)
	require.Equal(t, likesN, stats.Likes)
	require.Equal(t, commentsN, stats.Comments)
}
