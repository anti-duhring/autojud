package users

import (
	"regexp"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  *string   `json:"password"`
	CreatedAt string    `json:"created_at"`
	UpdatedAt string    `json:"updated_at"`
	DeletedAt *string   `json:"deleted_at"`
}

func (u *User) Validate() error {
	if u.Name == "" {
		return ErrNameEmpty
	}
	if u.Email == "" {
		return ErrEmailEmpty
	}
	if u.Password == nil || *u.Password == "" {
		return ErrPasswordEmpty
	}

	pattern := `^[a-zA-Z0-9_.+-]+@[a-zA-Z0-9-]+\.[a-zA-Z0-9-.]+$`
	if !regexp.MustCompile(pattern).MatchString(u.Email) {
		return ErrEmailInvalid
	}

	return nil
}
