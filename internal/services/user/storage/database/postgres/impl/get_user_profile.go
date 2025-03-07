package impl

import (
	"context"
	"errors"

	"github.com/LeonidS635/soa/internal/services/user/dto"
	"github.com/jackc/pgx/v5"
)

func (pg *PgUserStorageImpl) GetUserProfile(ctx context.Context, userId int) (*dto.Profile, error) {
	var profile dto.Profile
	err := pg.db.QueryRow(
		ctx,
		`SELECT name, surname, age, email, phone, city, description, birthdate, updated_at FROM profiles
        WHERE user_id = $1`,
		userId,
	).Scan(
		&profile.Name, &profile.Surname, &profile.Age, &profile.Email, &profile.Phone, &profile.City,
		&profile.Description, &profile.Birthdate, &profile.UpdatedAt,
	)

	if errors.Is(err, pgx.ErrNoRows) {
		err = UserNotFoundError
	}
	return &profile, err
}
