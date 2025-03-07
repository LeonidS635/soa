package dto

import (
	"crypto/md5"
	"time"
)

type RegistrationData struct {
	LoginData
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

type LoginData struct {
	Username string `json:"username"`
	Password string `json:"password"`

	PasswordHash [16]byte `json:"-"`
}

func (d *LoginData) ChangePasswordWithHash() {
	d.PasswordHash = md5.Sum([]byte(d.Password + d.Username))
}

type Profile struct {
	Name        string `json:"name"`
	Surname     string `json:"surname"`
	Age         int    `json:"age"`
	Email       string `json:"email"`
	Phone       string `json:"phone"`
	City        string `json:"city"`
	Description string `json:"description"`

	Birthdate time.Time `json:"birthdate"`
	UpdatedAt time.Time `json:"updated_at"`
}

//type dateJSON struct {
//	time.Time
//}
//
//func (d *dateJSON) MarshalJSON() ([]byte, error) {
//	log.Println(d)
//	log.Println(d.Format(time.DateOnly))
//	date := d.Format(time.DateOnly)
//	return []byte(date), nil
//}
//
//func (d dateJSON) UnmarshalJSON(b []byte) error {
//	date, err := time.Parse(time.DateOnly, string(b))
//	if err != nil {
//		return err
//	}
//
//	d.Time = date
//	return nil
//}
//
//func (d *dateJSON) Scan(value interface{}) error {
//	time_, ok := value.(time.Time)
//	if !ok {
//		return errors.New("type assertion time.Time failed")
//	}
//
//	d.Time = time_
//	return nil
//}
