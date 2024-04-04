// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package graphql

import (
	"fmt"
	"io"
	"strconv"
)

type AuthResponse struct {
	User     *User   `json:"user"`
	Token    string  `json:"token"`
	TokenExp float64 `json:"tokenExp"`
}

type CreateUserInput struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Mutation struct {
}

type Process struct {
	ID            string  `json:"id"`
	ProcessNumber string  `json:"processNumber"`
	Court         Court   `json:"court"`
	Origin        *string `json:"origin,omitempty"`
	Judge         *string `json:"judge,omitempty"`
	ActivePart    *string `json:"activePart,omitempty"`
	PassivePart   *string `json:"passivePart,omitempty"`
	CreatedAt     string  `json:"createdAt"`
	UpdatedAt     string  `json:"updatedAt"`
	DeletedAt     *string `json:"deletedAt,omitempty"`
}

type Query struct {
}

type UpdateUserInput struct {
	Name     *string `json:"name,omitempty"`
	Email    *string `json:"email,omitempty"`
	Password *string `json:"password,omitempty"`
}

type User struct {
	ID        string  `json:"id"`
	Name      string  `json:"name"`
	Email     string  `json:"email"`
	Password  *string `json:"password,omitempty"`
	CreatedAt string  `json:"createdAt"`
	UpdatedAt string  `json:"updatedAt"`
	DeletedAt *string `json:"deletedAt,omitempty"`
}

type Court string

const (
	CourtTjpe Court = "TJPE"
)

var AllCourt = []Court{
	CourtTjpe,
}

func (e Court) IsValid() bool {
	switch e {
	case CourtTjpe:
		return true
	}
	return false
}

func (e Court) String() string {
	return string(e)
}

func (e *Court) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Court(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Court", str)
	}
	return nil
}

func (e Court) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
