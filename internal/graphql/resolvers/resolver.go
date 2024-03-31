package resolvers

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

import "github.com/anti-duhring/autojud/internal/user"

type Resolver struct {
	UserService *user.Service
}
