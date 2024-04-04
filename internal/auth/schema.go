package auth

import "github.com/anti-duhring/autojud/internal/users"

type Response struct {
	User     users.User
	Token    string
	TokenExp int64
}
