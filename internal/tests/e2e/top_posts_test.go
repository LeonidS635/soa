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

func TestE2E_TopPosts(t *testing.T) {
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

	// Making new posts

	const postsN = 3
	postsIDs := make([]int, postsN)

	for i := 0; i < postsN; i++ {
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

		postsIDs[i] = postID.ID
	}

	// Like posts

	const likesN = 2
	for _, postID := range postsIDs {
		for i := 0; i < postID*likesN; i++ {
			likeResp, err := client.Post(baseURL+fmt.Sprintf("/posts/%d/like", postID), "", nil)
			if err != nil {
				t.Fatal(err)
			}
			_ = likeResp.Body.Close()

			require.Equal(t, http.StatusOK, likeResp.StatusCode)
		}
	}

	// Get top

	topPostsResp, err := client.Get(baseURL + "/posts/top?metric=likes")
	if err != nil {
		t.Fatal(err)
	}
	defer topPostsResp.Body.Close()

	b, _ := io.ReadAll(topPostsResp.Body)
	t.Log(string(b))

	require.Equal(t, http.StatusOK, topPostsResp.StatusCode)

	type post struct {
		PostID      int `json:"postId"`
		MetricValue int `json:"metricValue"`
	}
	var top []post
	require.NoError(t, json.NewDecoder(topPostsResp.Body).Decode(&top))

	require.Len(t, top, postsN)
	for i := postsN - 1; i >= 0; i-- {
		require.Equal(t, postsIDs[i], top[i].PostID)
		require.Equal(t, postsIDs[i]*likesN, top[i].MetricValue)
	}
}
