package api

import (
	"encoding/json"
	"github.com/gorilla/mux"
	httpApi "houseService/internal/handler/http"
	"houseService/internal/handler/http/middlewre"
	apimodel "houseService/internal/handler/http/model"
	"houseService/internal/service"
	"log/slog"
	"net/http"
)

type UserController struct {
	log     *slog.Logger
	service service.User
}

func NewUserController(log *slog.Logger, service service.User, router *mux.Router) *UserController {
	uc := &UserController{
		log:     log,
		service: service,
	}
	router.HandleFunc("/register", uc.RegisterPost).Methods(http.MethodPost)
	router.HandleFunc("/login", uc.LoginPost).Methods(http.MethodPost)

	return uc
}

func (u *UserController) LoginPost(w http.ResponseWriter, r *http.Request) {
	const op = "handler.http.api.LoginPost"

	log := u.log.With(
		slog.String("op", op),
		slog.String("request_id", r.Context().Value(middlewre.CtxKeyRequestID).(string)),
	)
	req := &apimodel.LoginBody{}

	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		log.Error("failed to decode request")
		ErrorResponseWithCode(w, r, http.StatusBadRequest, err)
		return
	}

	token, err := u.service.SingIn(httpApi.ToUserLoginServiceModel(req))
	if err != nil {
		log.Error("failed to login user", slog.String("error", err.Error()))
		ErrorResponse(w, r, err)
		return
	}

	res := httpApi.ToTokenApiModel(token)

	response(w, r, http.StatusOK, res)
}

func (u *UserController) RegisterPost(w http.ResponseWriter, r *http.Request) {
	const op = "handler.http.api.RegisterPost"

	log := u.log.With(
		slog.String("op", op),
		slog.String("request_id", r.Context().Value(middlewre.CtxKeyRequestID).(string)),
	)
	req := &apimodel.RegisterBody{}

	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		log.Error("failed to decode request")
		ErrorResponseWithCode(w, r, http.StatusBadRequest, err)
		return
	}

	userServiceModel, err := httpApi.ToUserServiceModel(req)
	if err != nil {
		log.Error("failed to decode request")
		ErrorResponseWithCode(w, r, http.StatusBadRequest, err)
		return
	}

	uuid, err := u.service.SingUp(userServiceModel)
	if err != nil {
		log.Error("failed to register user", slog.String("error", err.Error()))
		ErrorResponse(w, r, err)
		return
	}

	log.Info("create new user")

	res := httpApi.ToUUIDApiModel(uuid)

	response(w, r, http.StatusOK, res)
}
