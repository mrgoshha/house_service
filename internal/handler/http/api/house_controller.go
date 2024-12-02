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

type HouseController struct {
	log         *slog.Logger
	service     service.House
	authManager *AuthManager
}

func NewHouseController(log *slog.Logger, service service.House, router *mux.Router, authManager *AuthManager) *HouseController {
	hc := &HouseController{
		log:         log,
		service:     service,
		authManager: authManager,
	}
	onlyModerator := router.PathPrefix("").Subrouter()
	onlyModerator.Use(authManager.UserIdentity)
	onlyModerator.Use(authManager.OnlyModerator)
	onlyModerator.HandleFunc("/house/create", hc.houseCreatePost)

	return hc
}

func (h *HouseController) houseCreatePost(w http.ResponseWriter, r *http.Request) {
	const op = "handler.http.api.HouseCreatePost"

	log := h.log.With(
		slog.String("op", op),
		slog.String("request_id", r.Context().Value(middlewre.CtxKeyRequestID).(string)),
	)

	req := &apimodel.HouseCreateBody{}

	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		log.Error("failed to decode request")
		ErrorResponseWithCode(w, r, http.StatusBadRequest, err)
		return
	}

	house, err := h.service.HouseCreate(httpApi.ToHouseServiceCreateModel(req))
	if err != nil {
		log.Error("failed to create house", slog.String("error", err.Error()))
		ErrorResponse(w, r, err)
		return
	}

	log.Info("house created")

	res := httpApi.ToHouseApiModel(house)

	response(w, r, http.StatusOK, res)
}
