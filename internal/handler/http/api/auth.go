package api

import (
	"context"
	"fmt"
	"houseService/internal/model"
	"houseService/pkg/auth"
	"net/http"
	"strings"
)

const (
	AuthorizationHeader = "Authorization"
	CtxKeyUserId        = "userId"
	CtxKeyUserType      = "userType"
)

type AuthManager struct {
	tokenManager auth.TokenManager
}

func NewTokenManager(t auth.TokenManager) *AuthManager {
	return &AuthManager{
		tokenManager: t,
	}
}

func (t *AuthManager) UserIdentity(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get(AuthorizationHeader)
		if header == "" {
			ErrorResponseWithCode(w, r, http.StatusUnauthorized, fmt.Errorf("empty auth header"))
			return
		}

		headerParts := strings.Split(header, " ")
		if len(headerParts) != 2 || headerParts[0] != "Bearer" {
			ErrorResponseWithCode(w, r, http.StatusUnauthorized, fmt.Errorf("invalid auth header"))
			return
		}

		if len(headerParts[1]) == 0 {
			ErrorResponseWithCode(w, r, http.StatusUnauthorized, fmt.Errorf("token is empty"))
			return
		}

		userId, userType, err := t.tokenManager.Parse(headerParts[1])
		if err != nil {
			ErrorResponseWithCode(w, r, http.StatusUnauthorized, err)
			return
		}

		ctx := context.WithValue(r.Context(), CtxKeyUserId, userId)
		ctx = context.WithValue(ctx, CtxKeyUserType, userType)

		next.ServeHTTP(w, r.WithContext(ctx))

	})
}

func (t *AuthManager) OnlyModerator(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ut := r.Context().Value(CtxKeyUserType)

		if ut != string(model.MODERATOR) {
			ErrorResponseWithCode(w, r, http.StatusForbidden, fmt.Errorf("access denied"))
			return
		}

		next.ServeHTTP(w, r)
	})

}
