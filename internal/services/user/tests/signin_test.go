package tests

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/LeonidS635/soa/internal/services/user/dto"
	"github.com/stretchr/testify/require"
)

func TestUserService_SignIn(t *testing.T) {
	t.Cleanup(
		func() {
			if err := clearTables(context.Background()); err != nil {
				t.Fatal(err)
			}
		},
	)

	const usersN = 10
	userIDs := make([]int, usersN)
	for i := 0; i < usersN; i++ {
		loginData := dto.RegistrationData{
			LoginData: dto.LoginData{
				Username: fmt.Sprintf("test %d", i),
				Password: fmt.Sprintf("qwerty#%d", i),
			},
			Email:     "test@email.com",
			CreatedAt: time.Now(),
		}
		loginData.ChangePasswordWithHash()
		profile := dto.Profile{
			Name:        fmt.Sprintf("Name %d", i),
			Surname:     fmt.Sprintf("Surname %d", i),
			Age:         5 * i,
			Email:       "test@email.com",
			Phone:       "1234567890",
			City:        fmt.Sprintf("City %d", i),
			Description: fmt.Sprintf("Description %d", i),
			Birthdate:   time.Now(),
			UpdatedAt:   time.Now(),
		}

		var userID int
		var err error
		if userID, err = addUser(t.Context(), loginData, profile); err != nil {
			t.Fatal(err)
		}
		userIDs[i] = userID
	}

	for i := 0; i < usersN; i++ {
		loginData := dto.LoginData{
			Username: fmt.Sprintf("test %d", i),
			Password: fmt.Sprintf("qwerty#%d", i),
		}

		body, err := json.Marshal(loginData)
		if err != nil {
			t.Fatal(err)
		}
		req := httptest.NewRequestWithContext(t.Context(), "POST", "/signin", bytes.NewReader(body))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		serverRouter.ServeHTTP(w, req)

		resp := w.Result()
		defer resp.Body.Close()

		require.Equal(t, http.StatusOK, resp.StatusCode)

		userID, err := getUserIDFromJWT(resp.Cookies())
		require.NoError(t, err)

		require.Equal(t, userIDs[i], userID)
	}
}
