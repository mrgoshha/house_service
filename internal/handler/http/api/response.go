package api

import (
	"encoding/json"
	"errors"
	"houseService/internal/adapter/dbs"
	"houseService/internal/handler/http/middlewre"
	"houseService/internal/handler/http/model"
	"houseService/internal/service"
	"net/http"
)

func ErrorResponseWithCode(w http.ResponseWriter, r *http.Request, code int, err error) {
	var modelErr interface{}
	if code == http.StatusInternalServerError {
		modelErr = model.InlineResponse500{
			RequestId: r.Context().Value(middlewre.CtxKeyRequestID).(string),
		}
	} else {
		modelErr = model.InlineResponseError{
			Message:   err.Error(),
			RequestId: r.Context().Value(middlewre.CtxKeyRequestID).(string),
		}
	}
	response(w, r, code, modelErr)

}

func ErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	var code int

	switch {
	case errors.Is(err, dbs.ErrorRecordNotFound):
		code = http.StatusNotFound
	case errors.Is(err, dbs.ErrorRecordAlreadyExists):
		code = http.StatusConflict
	case errors.Is(err, service.Invalid):
		code = http.StatusBadRequest
	default:
		code = http.StatusInternalServerError
	}
	ErrorResponseWithCode(w, r, code, err)
}

func response(w http.ResponseWriter, _ *http.Request, code int, data interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(code)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}
