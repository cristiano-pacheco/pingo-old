// Package authmw contains the middleware to handle the authentication via token.
package authmw

import (
	"context"
	"net/http"

	"github.com/cristiano-pacheco/pingo/internal/application/service/tokensvc"
	"github.com/cristiano-pacheco/pingo/internal/infra/http/request"
	"github.com/cristiano-pacheco/pingo/internal/infra/http/response"
)

type AuthMiddleware struct {
	tokenService tokensvc.TokenService
}

func New(tokenService tokensvc.TokenService) *AuthMiddleware {
	return &AuthMiddleware{tokenService: tokenService}
}

func (m *AuthMiddleware) Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		bearerToken := r.Header.Get("Authorization")
		claims, err := m.tokenService.ParseToken(r.Context(), bearerToken)
		if err != nil {
			response.UnauthorizedResponse(w, r)
			return
		}
		ctx := context.WithValue(r.Context(), request.UserIDContextKey, claims.Subject)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
