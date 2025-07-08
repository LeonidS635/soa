package tests

import (
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

func TestUserService_GetProfile(t *testing.T) {
	t.Cleanup(
		func() {
			if err := clearTables(context.Background()); err != nil {
				t.Fatal(err)
			}
		},
	)

	const usersN = 10
	userIDs := make([]int, usersN)
	profiles := make([]dto.Profile, usersN)
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
		profiles[i] = dto.Profile{
			Name:        fmt.Sprintf("Name %d", i),
			Surname:     fmt.Sprintf("Surname %d", i),
			Age:         5 * i,
			Email:       "test@email.com",
			Phone:       "1234567890",
			City:        fmt.Sprintf("City %d", i),
			Description: fmt.Sprintf("Description %d", i),
			Birthdate:   time.Date(2000, 01, 01, 00, 00, 00, 00, time.Local),
			UpdatedAt:   time.Now(),
		}

		var userID int
		var err error
		if userID, err = addUser(t.Context(), loginData, profiles[i]); err != nil {
			t.Fatal(err)
		}
		userIDs[i] = userID
	}

	for i := 0; i < usersN; i++ {
		jwtToken, err := generateJWT(userIDs[i])
		if err != nil {
			t.Fatal(err)
		}

		req := httptest.NewRequestWithContext(t.Context(), "GET", "/profile", nil)
		req.AddCookie(&http.Cookie{Name: "jwt", Value: jwtToken})

		w := httptest.NewRecorder()
		serverRouter.ServeHTTP(w, req)

		resp := w.Result()
		defer resp.Body.Close()

		respProfile := dto.Profile{}
		err = json.NewDecoder(resp.Body).Decode(&respProfile)
		require.NoError(t, err)

		profiles[i].UpdatedAt = respProfile.UpdatedAt

		require.Equal(t, http.StatusOK, resp.StatusCode)
		require.Equal(t, profiles[i], respProfile)
	}
}
