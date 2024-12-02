package service

import (
	"fmt"
	"os"
	"time"
)

const (
	SigningKey     = "JWT_SIGNING_KEY"
	AccessTokenTTL = "ACCESS_TOKEN_TTL"
)

type Config struct {
	SigningKey     string
	AccessTokenTTL time.Duration
}

func NewConfig() (*Config, error) {
	key := os.Getenv(SigningKey)
	if len(key) == 0 {
		return nil, fmt.Errorf("empty signing key")
	}
	aTtl := os.Getenv(AccessTokenTTL)
	if len(aTtl) == 0 {
		aTtl = "2h"
	}

	at, _ := time.ParseDuration(aTtl)
	return &Config{
		SigningKey:     SigningKey,
		AccessTokenTTL: at,
	}, nil
}
