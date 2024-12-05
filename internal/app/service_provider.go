package app

import (
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	flatRepo "houseService/internal/adapter/dbs/postrges/flat"
	houseRepo "houseService/internal/adapter/dbs/postrges/house"
	userRepo "houseService/internal/adapter/dbs/postrges/user"
	"houseService/internal/handler/http"
	"houseService/internal/handler/http/api"
	"houseService/internal/service/flat"
	"houseService/internal/service/house"
	"houseService/internal/service/user"
	"houseService/pkg/auth"
	"houseService/pkg/hash"
	"log/slog"
)

type ServiceProvider struct {
	houseController *api.HouseController
	flatController  *api.FlatController
	userController  *api.UserController
	houseService    *house.Service
	flatService     *flat.Service
	userService     *user.Service
	houseRepository *houseRepo.Repository
	flatRepository  *flatRepo.Repository
	userRepository  *userRepo.Repository
	httpRouter      *mux.Router
	log             *slog.Logger
	db              *sqlx.DB
	hash            *hash.SHA1Hasher
	tokenManager    *auth.Manager
	auth            *api.AuthManager
}

func NewServiceProvider(log *slog.Logger, db *sqlx.DB, t *auth.Manager) *ServiceProvider {
	return &ServiceProvider{
		log:          log,
		db:           db,
		tokenManager: t,
	}
}

func (s *ServiceProvider) HouseController() *api.HouseController {
	if s.houseController == nil {
		s.houseController = api.NewHouseController(
			s.log, s.HouseService(), s.HttpRouter(), s.Auth(),
		)
	}
	return s.houseController
}

func (s *ServiceProvider) FlatController() *api.FlatController {
	if s.flatController == nil {
		s.flatController = api.NewFlatController(
			s.log, s.FlatService(), s.HttpRouter(), s.Auth(),
		)
	}
	return s.flatController
}

func (s *ServiceProvider) UserController() *api.UserController {
	if s.userController == nil {
		s.userController = api.NewUserController(
			s.log, s.UserService(), s.HttpRouter(),
		)
	}
	return s.userController
}

func (s *ServiceProvider) HouseService() *house.Service {
	if s.houseService == nil {
		s.houseService = house.NewService(
			s.HouseRepository(),
		)
	}
	return s.houseService
}

func (s *ServiceProvider) FlatService() *flat.Service {
	if s.flatService == nil {
		s.flatService = flat.NewService(
			s.FlatRepository(), s.HouseService(),
		)
	}
	return s.flatService
}

func (s *ServiceProvider) UserService() *user.Service {
	if s.userService == nil {
		s.userService = user.NewService(
			s.UserRepository(), s.Hash(), s.tokenManager,
		)
	}
	return s.userService
}

func (s *ServiceProvider) HouseRepository() *houseRepo.Repository {
	if s.houseRepository == nil {
		s.houseRepository = houseRepo.NewRepository(
			s.db,
		)
	}
	return s.houseRepository
}

func (s *ServiceProvider) FlatRepository() *flatRepo.Repository {
	if s.flatRepository == nil {
		s.flatRepository = flatRepo.NewRepository(
			s.db,
		)
	}
	return s.flatRepository
}

func (s *ServiceProvider) UserRepository() *userRepo.Repository {
	if s.userRepository == nil {
		s.userRepository = userRepo.NewRepository(
			s.db,
		)
	}
	return s.userRepository
}

func (s *ServiceProvider) HttpRouter() *mux.Router {
	if s.httpRouter == nil {
		s.httpRouter = http.NewRouter(s.log)
	}
	return s.httpRouter
}

func (s *ServiceProvider) RegisterControllers() {
	if s.houseController == nil {
		s.HouseController()
		s.FlatController()
		s.UserController()
	}
}

func (s *ServiceProvider) Hash() *hash.SHA1Hasher {
	if s.hash == nil {
		s.hash = hash.NewSHA1Hasher()
	}
	return s.hash
}

func (s *ServiceProvider) Auth() *api.AuthManager {
	if s.auth == nil {
		s.auth = api.NewTokenManager(s.tokenManager)
	}
	return s.auth
}
