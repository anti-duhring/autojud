package auth

import "context"

var userIdKey = &contextKey{"userKey"}

var tokenKey = &contextKey{"tokenKey"}

type contextKey struct {
	name string
}

func SaveUserID(ctx context.Context, userID string) context.Context {
	return context.WithValue(ctx, userIdKey, userID)
}

func GetUserID(ctx context.Context) string {
	raw, _ := ctx.Value(userIdKey).(string)
	return raw
}

func SaveToken(ctx context.Context, token string) context.Context {
	return context.WithValue(ctx, tokenKey, token)
}

func GetToken(ctx context.Context) string {
	raw, _ := ctx.Value(tokenKey).(string)
	return raw
}
