package tests

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUserService_SignUp(t *testing.T) {
	t.Cleanup(
		func() {
			if err := clearTables(context.Background()); err != nil {
				t.Fatal(err)
			}
		},
	)

	const usersN = 10
	for i := 0; i < usersN; i++ {
		registrationData := struct {
			Username string `json:"username"`
			Password string `json:"password"`
			Email    string `json:"email"`
		}{
			Username: fmt.Sprintf("john_doe_%d", i),
			Password: "securePassword123",
			Email:    fmt.Sprintf("john%d@example.com", i),
		}

		body, err := json.Marshal(registrationData)
		if err != nil {
			t.Fatal(err)
		}
		req := httptest.NewRequestWithContext(t.Context(), "POST", "/signup", bytes.NewReader(body))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		serverRouter.ServeHTTP(w, req)

		resp := w.Result()
		defer resp.Body.Close()

		require.Equal(t, http.StatusCreated, resp.StatusCode)

		userID, err := getUserIDFromJWT(resp.Cookies())
		require.NoError(t, err)
		require.Positive(t, userID)

		var username, email string
		err = db.QueryRow(
			t.Context(),
			`SELECT username, email FROM users WHERE user_id = $1`,
			userID,
		).Scan(&username, &email)
		require.NoError(t, err)

		require.Equal(t, registrationData.Username, username)
		require.Equal(t, registrationData.Email, email)
	}
}
