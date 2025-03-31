package impl

import (
	"context"
	"errors"

	"github.com/LeonidS635/soa/internal/services/user/dto"
)

func (pg *PgUserStorageImpl) AddUser(ctx context.Context, credentials *dto.RegistrationData) (int, error) {
	_, err := pg.CheckUser(ctx, &credentials.LoginData)
	if err == nil {
		return -1, UserAlreadyExistsError
	}
	if !errors.Is(err, UserNotFoundError) {
		return -1, err
	}

	var userId int
	err = pg.db.QueryRow(
		ctx,
		`INSERT INTO users (username, password_hash, email) VALUES ($1, $2, $3) RETURNING user_id`,
		credentials.Username, credentials.PasswordHash[:], credentials.Email,
	).Scan(&userId)
	if err != nil {
		return -1, err
	}

	_, err = pg.db.Exec(
		ctx,
		`INSERT INTO profiles (user_id, email) VALUES ($1, $2)`,
		userId, credentials.Email,
	)
	return userId, err
}
