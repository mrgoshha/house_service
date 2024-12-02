package api

import (
	"encoding/json"
	"github.com/gorilla/mux"
	httpApi "houseService/internal/handler/http"
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
	req := &apimodel.LoginBody{}

	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		ErrorResponseWithCode(w, r, http.StatusBadRequest, err)
		return
	}

	token, err := u.service.SingIn(httpApi.ToUserLoginServiceModel(req))
	if err != nil {
		ErrorResponse(w, r, err)
		return
	}

	res := httpApi.ToTokenApiModel(token)

	response(w, r, http.StatusOK, res)
}

func (u *UserController) RegisterPost(w http.ResponseWriter, r *http.Request) {
	req := &apimodel.RegisterBody{}

	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		ErrorResponseWithCode(w, r, http.StatusBadRequest, err)
		return
	}

	userServiceModel, err := httpApi.ToUserServiceModel(req)
	if err != nil {
		ErrorResponseWithCode(w, r, http.StatusBadRequest, err)
		return
	}

	uuid, err := u.service.SingUp(userServiceModel)
	if err != nil {
		ErrorResponse(w, r, err)
		return
	}

	res := httpApi.ToUUIDApiModel(uuid)

	response(w, r, http.StatusOK, res)
}
