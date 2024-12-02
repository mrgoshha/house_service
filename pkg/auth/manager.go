package auth

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type TokenManager interface {
	NewJWT(userId, userType string) (string, error)
	Parse(accessToken string) (string, string, error)
}

type Manager struct {
	signingKey     string
	accessTokenTTL time.Duration
}

type tokenClaims struct {
	jwt.StandardClaims
	UserType string `json:"user_type"`
}

func NewManager(signingKey string, accessTokenTTL time.Duration) (*Manager, error) {
	if signingKey == "" {
		return nil, errors.New("empty signing key")
	}

	return &Manager{
		signingKey:     signingKey,
		accessTokenTTL: accessTokenTTL,
	}, nil
}

func (m *Manager) NewJWT(userId, userType string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(m.accessTokenTTL).Unix(),
			Subject:   userId,
		},
		UserType: userType,
	})

	return token.SignedString([]byte(m.signingKey))
}

func (m *Manager) Parse(accessToken string) (string, string, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(m.signingKey), nil
	})
	if err != nil {
		return "", "", err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return "", "", fmt.Errorf("token claims are not of type *tokenClaims")
	}

	return claims.Subject, claims.UserType, nil
}
