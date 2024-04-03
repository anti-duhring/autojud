package auth

import (
	"net/http"

	"github.com/anti-duhring/autojud/internal/user"
	"github.com/anti-duhring/autojud/pkg/jwt"
	"github.com/anti-duhring/goncurrency/pkg/logger"
	"github.com/google/uuid"
)

func Middleware(userService user.Service) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			header := r.Header.Get("Authorization")

			ctx := r.Context()
			r = r.WithContext(ctx)

			// Allow unauthenticated users in
			if header == "" {
				next.ServeHTTP(w, r)
				return
			}

			//validate jwt token
			tokenStr := header
			jwtClaims, err := jwt.ParseToken(tokenStr)
			if err != nil {
				logger.Error("error parsing token", err)
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
			ctx = SaveUserID(ctx, jwtClaims.UserID)
			ctx = SaveToken(ctx, tokenStr)

			// and call the next with our new context
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}
