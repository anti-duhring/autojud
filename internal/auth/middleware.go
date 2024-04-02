package auth

import (
	"context"
	"net/http"

	"github.com/anti-duhring/autojud/internal/user"
	"github.com/anti-duhring/autojud/pkg/jwt"
	"github.com/google/uuid"
)

func Middleware(userService user.Service) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			header := r.Header.Get("Authorization")

			// Allow unauthenticated users in
			if header == "" {
				next.ServeHTTP(w, r)
				return
			}

			//validate jwt token
			tokenStr := header
			jwtClaims, err := jwt.ParseToken(tokenStr)
			if err != nil {
				http.Error(w, "Invalid token", http.StatusForbidden)
				return
			}

			if jwt.TokenExpired(int64(jwtClaims.Exp)) {
				http.Error(w, "Token expired", http.StatusForbidden)
				return
			}

			// create user and check if user exists in db
			userIDStr, err := uuid.Parse(jwtClaims.UserID)
			if err != nil {
				http.Error(w, "Invalid token", http.StatusForbidden)
				return
			}

			_, err = userService.GetByID(userIDStr, r.Context())
			if err != nil {
				next.ServeHTTP(w, r)
				return
			}
			// put it in context
			ctx := context.Background()
			SaveUserID(ctx, jwtClaims.UserID)
			SaveToken(ctx, tokenStr)

			// and call the next with our new context
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}
