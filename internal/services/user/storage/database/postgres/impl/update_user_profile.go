package impl

import (
	"context"
	"errors"

	"github.com/LeonidS635/soa/internal/services/user/dto"
	"github.com/jackc/pgx/v5"
)

func (pg *PgUserStorageImpl) UpdateUserProfile(ctx context.Context, userId int, newProfile *dto.Profile) error {
	_, err := pg.db.Exec(
		ctx,
		`UPDATE profiles
		SET name=$1, surname=$2, age=$3, email=$4, phone=$5, city=$6, description=$7, birthdate=$8
		WHERE user_id=$9`,
		newProfile.Name, newProfile.Surname, newProfile.Age, newProfile.Email, newProfile.Phone, newProfile.City,
		newProfile.Description, newProfile.Birthdate, userId,
	)

	if errors.Is(err, pgx.ErrNoRows) {
		err = UserNotFoundError
	}
	return err
}
