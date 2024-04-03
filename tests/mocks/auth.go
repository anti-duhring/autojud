package mocks

import (
	"github.com/99designs/gqlgen/client"
	"github.com/anti-duhring/autojud/internal/auth"
	"github.com/anti-duhring/autojud/pkg/jwt"
)

func MockToken() client.Option {
	return func(bd *client.Request) {
		tokenStr, _, _ := jwt.GenerateToken("e7144ad8-6a6f-446a-b735-c96063a314e0")
		ctx := auth.SaveToken(bd.HTTP.Context(), tokenStr)
		bd.HTTP = bd.HTTP.WithContext(ctx)
	}
}

func MockUserID(id string) client.Option {
	return func(bd *client.Request) {
		ctx := auth.SaveUserID(bd.HTTP.Context(), id)
		bd.HTTP = bd.HTTP.WithContext(ctx)
	}
}
