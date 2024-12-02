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
	"strconv"
)

type FlatController struct {
	log         *slog.Logger
	service     service.Flat
	authManager *AuthManager
}

func NewFlatController(log *slog.Logger, service service.Flat, router *mux.Router, authManager *AuthManager) *FlatController {
	fc := &FlatController{
		log:         log,
		service:     service,
		authManager: authManager,
	}
	authRouter := router.PathPrefix("").Subrouter()
	authRouter.Use(authManager.UserIdentity)
	authRouter.HandleFunc("/flat/create", fc.FlatCreatePost).Methods(http.MethodPost)
	authRouter.HandleFunc("/house", fc.HouseIdGet).Methods(http.MethodGet)

	onlyModerator := router.PathPrefix("").Subrouter()
	onlyModerator.Use(authManager.UserIdentity)
	onlyModerator.Use(authManager.OnlyModerator)
	onlyModerator.HandleFunc("/flat/update", fc.FlatUpdatePost).Methods(http.MethodPost)

	return fc
}

func (f *FlatController) FlatCreatePost(w http.ResponseWriter, r *http.Request) {
	const op = "handler.http.api.FlatCreatePost"

	log := f.log.With(
		slog.String("op", op),
		slog.String("request_id", r.Context().Value(middlewre.CtxKeyRequestID).(string)),
	)
	req := &apimodel.FlatCreateBody{}

	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		log.Error("failed to decode request")
		ErrorResponseWithCode(w, r, http.StatusBadRequest, err)
		return
	}

	flat, err := f.service.FlatCreate(httpApi.ToFlatServiceModel(req))
	if err != nil {
		log.Error("failed to create flat", slog.String("error", err.Error()))
		ErrorResponse(w, r, err)
		return
	}

	log.Info("flat created", slog.Int("id", flat.Id))

	res := httpApi.ToFlatApiModel(flat)

	response(w, r, http.StatusOK, res)
}

func (f *FlatController) FlatUpdatePost(w http.ResponseWriter, r *http.Request) {
	const op = "handler.http.api.FlatUpdatePost"

	log := f.log.With(
		slog.String("op", op),
		slog.String("request_id", r.Context().Value(middlewre.CtxKeyRequestID).(string)),
	)
	req := &apimodel.FlatUpdateBody{}

	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		log.Error("failed to decode request")
		ErrorResponseWithCode(w, r, http.StatusBadRequest, err)
		return
	}

	flatServiceModel, err := httpApi.ToFlatUpdateServiceModel(req)
	if err != nil {
		log.Error("failed to convert model", slog.String("error", err.Error()))
		ErrorResponseWithCode(w, r, http.StatusBadRequest, err)
		return
	}

	flat, err := f.service.FlatUpdate(flatServiceModel)
	if err != nil {
		log.Error("failed to update flat", slog.String("error", err.Error()))
		ErrorResponse(w, r, err)
		return
	}

	log.Info("flat updated")

	res := httpApi.ToFlatApiModel(flat)

	response(w, r, http.StatusOK, res)
}

func (f *FlatController) HouseIdGet(w http.ResponseWriter, r *http.Request) {
	const op = "handler.http.api.HouseIdGet"

	log := f.log.With(
		slog.String("op", op),
		slog.String("request_id", r.Context().Value(middlewre.CtxKeyRequestID).(string)),
	)
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		log.Error("failed to decode request")
		ErrorResponseWithCode(w, r, http.StatusBadRequest, err)
		return
	}

	flats, err := f.service.HouseFlatsGet(id, r.Context().Value(CtxKeyUserType).(string))
	if err != nil {
		log.Error("failed to get flats", slog.String("error", err.Error()))
		ErrorResponse(w, r, err)
		return
	}

	res := httpApi.ToFlatsApiModel(flats)

	response(w, r, http.StatusOK, res)
}
