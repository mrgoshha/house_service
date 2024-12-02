package user

import (
	"fmt"
	"github.com/google/uuid"
	"houseService/internal/adapter/dbs"
	"houseService/internal/model"
	"houseService/internal/service"
	"houseService/pkg/auth"
	"houseService/pkg/hash"
)

type Service struct {
	repository   dbs.UserRepository
	hasher       hash.PasswordHasher
	tokenManager auth.TokenManager
}

func NewService(r dbs.UserRepository, h hash.PasswordHasher, t auth.TokenManager) *Service {
	return &Service{
		repository:   r,
		hasher:       h,
		tokenManager: t,
	}
}

func (s *Service) SingUp(u *model.User) (string, error) {
	id := uuid.New().String()
	passwordHash, err := s.hasher.Hash(u.Password)
	if err != nil {
		return "", fmt.Errorf(`hash password %w`, err)
	}

	user := &model.User{
		Id:       id,
		Email:    u.Email,
		Password: u.Password,
		UserType: u.UserType,
	}

	if err = user.Validate(); err != nil {
		return "", fmt.Errorf(`validate user %w`, service.Invalid)
	}

	user.Password = passwordHash

	if err = s.repository.CreateUser(user); err != nil {
		return "", err
	}

	return id, nil
}

func (s *Service) SingIn(u *model.UserLogin) (string, error) {
	passwordHash, err := s.hasher.Hash(u.Password)
	if err != nil {
		return "", fmt.Errorf(`hash password %w`, err)
	}

	u.Password = passwordHash

	user, err := s.repository.GetUserByCredentials(u)
	if err != nil {
		return "", err
	}

	return s.createToken(user)
}

func (s *Service) createToken(u *model.User) (string, error) {
	jwt, err := s.tokenManager.NewJWT(u.Id, string(u.UserType))
	if err != nil {
		return "", fmt.Errorf(`create token %w`, err)
	}
	return jwt, nil
}
