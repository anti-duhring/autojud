package auth

import "github.com/anti-duhring/autojud/internal/user"

type Response struct {
	User     user.User
	Token    string
	TokenExp int64
}
