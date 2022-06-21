package middleware

import (
	"context"
	"net/http"

	"github.com/tmammado/take-home-assignment/jwt"
)

func JwtAuthentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		tokenString := req.Header.Get("x-authentication-token")
		claims, err := jwt.Verify(tokenString)
		if err != nil {
			res.WriteHeader(http.StatusUnauthorized)
			return
		}
		ctx := context.WithValue(req.Context(), "email", claims.Email)
		req = req.WithContext(ctx)
		next.ServeHTTP(res, req)
	})
}
