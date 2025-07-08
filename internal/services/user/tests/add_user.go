package tests

import (
	"context"

	"github.com/LeonidS635/soa/internal/services/user/dto"
)

func addUser(ctx context.Context, loginData dto.RegistrationData, profile dto.Profile) (int, error) {
	var userID int

	err := db.QueryRow(
		ctx,
		`INSERT INTO users (username, password_hash, email, created_at) VALUES ($1, $2, $3, $4) RETURNING user_id`,
		loginData.Username, loginData.PasswordHash[:], loginData.Email, loginData.CreatedAt,
	).Scan(&userID)
	if err != nil {
		return -1, err
	}

	_, err = db.Exec(
		ctx,
		`INSERT INTO profiles VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`,
		userID, profile.Name, profile.Surname, profile.Age, profile.Email, profile.Phone, profile.City,
		profile.Description, profile.Birthdate, loginData.CreatedAt,
	)
	if err != nil {
		return -1, err
	}

	return userID, nil
}
