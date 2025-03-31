package impl

import (
	"context"
	"errors"

	"github.com/LeonidS635/soa/internal/services/user/dto"
	"github.com/jackc/pgx/v5"
)

func (pg *PgUserStorageImpl) CheckUser(ctx context.Context, user *dto.LoginData) (int, error) {
	var userId int
	var passwordHash []byte
	err := pg.db.QueryRow(
		ctx,
		`SELECT user_id, password_hash FROM users WHERE username = $1`,
		user.Username,
	).Scan(&userId, &passwordHash)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			err = UserNotFoundError
		}
		return -1, err
	}

	var passwordHash16Bytes [16]byte
	if len(passwordHash) != 16 {
		return -1, errors.New("saved password is corrupted")
	}
	copy(passwordHash16Bytes[:], passwordHash)

	if passwordHash16Bytes != user.PasswordHash {
		return -1, IncorrectPasswordError
	}
	return userId, nil
}
