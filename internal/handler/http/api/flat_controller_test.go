package api

import (
	"bytes"
	"context"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	httpApi "houseService/internal/handler/http"
	"houseService/internal/model"
	servicemocks "houseService/internal/service/mocks"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestFlatController_FlatCreatePost(t *testing.T) {
	// Init Test Table
	type mockBehavior func(r *servicemocks.MockFlat, user *model.Flat)

	tests := []struct {
		name                 string
		inputBody            string
		inputFlat            *model.Flat
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name: "Ok",
			inputBody: `{
				"house_id": 1,
				"price": 20,
				"rooms": 2
			}`,
			inputFlat: &model.Flat{
				HouseId: 1,
				Price:   20,
				Rooms:   2,
			},
			mockBehavior: func(r *servicemocks.MockFlat, flat *model.Flat) {
				newFlat := &model.Flat{
					Id:      1,
					HouseId: 1,
					Price:   20,
					Rooms:   2,
					Status:  model.CREATED,
				}
				r.EXPECT().FlatCreate(flat).Return(newFlat, nil)
			},
			expectedStatusCode: http.StatusOK,
			expectedResponseBody: `{"id":1,"house_id":1,"price":20,"rooms":2,"status":"created"}
`,
		},
		{
			name:               "Wrong Input",
			inputBody:          `invalid`,
			inputFlat:          nil,
			mockBehavior:       func(r *servicemocks.MockFlat, flat *model.Flat) {},
			expectedStatusCode: http.StatusBadRequest,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Init Dependencies
			c := gomock.NewController(t)
			defer c.Finish()

			service := servicemocks.NewMockFlat(c)
			test.mockBehavior(service, test.inputFlat)

			log := slog.New(
				slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))

			controller := &FlatController{
				service: service,
				log:     log,
			}

			// Init Endpoint
			r := httpApi.NewRouter(log)
			r.HandleFunc("/flat/create", controller.FlatCreatePost).Methods(http.MethodPost)

			// Create Request
			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/flat/create",
				bytes.NewBufferString(test.inputBody))

			// Make Request
			r.ServeHTTP(w, req)

			// Assert
			assert.Equal(t, w.Code, test.expectedStatusCode)
			if w.Code == http.StatusOK {
				assert.Equal(t, w.Body.String(), test.expectedResponseBody)
			}
		})
	}
}

func TestFlatController_HouseIdGet(t *testing.T) {
	// Init Test Table
	type mockBehavior func(r *servicemocks.MockFlat, id int, userType string)

	tests := []struct {
		name                 string
		inputId              int
		url                  string
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:    "Ok",
			inputId: 1,
			url:     "/house?id=1",
			mockBehavior: func(r *servicemocks.MockFlat, id int, userType string) {
				flats := []*model.Flat{
					{
						Id:      1,
						HouseId: 1,
						Price:   20,
						Rooms:   2,
						Status:  model.APPROVED,
					},
				}
				r.EXPECT().HouseFlatsGet(id, userType).Return(flats, nil)
			},
			expectedStatusCode: http.StatusOK,
			expectedResponseBody: `{"flats":[{"id":1,"house_id":1,"price":20,"rooms":2,"status":"approved"}]}
`,
		},
		{
			name:               "Wrong Input",
			inputId:            1,
			url:                "/house?id=1r",
			mockBehavior:       func(r *servicemocks.MockFlat, id int, userType string) {},
			expectedStatusCode: http.StatusBadRequest,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Init Dependencies
			c := gomock.NewController(t)
			defer c.Finish()

			service := servicemocks.NewMockFlat(c)
			test.mockBehavior(service, test.inputId, "")

			log := slog.New(
				slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))

			controller := &FlatController{
				service: service,
				log:     log,
			}

			// Init Endpoint
			r := httpApi.NewRouter(log)
			r.HandleFunc("/house", controller.HouseIdGet).Methods(http.MethodGet)

			// Create Request
			w := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodGet, test.url, nil)
			ctx := context.WithValue(req.Context(), CtxKeyUserType, "")

			// Make Request
			r.ServeHTTP(w, req.WithContext(ctx))

			// Assert
			assert.Equal(t, w.Code, test.expectedStatusCode)
			if w.Code == http.StatusOK {
				assert.Equal(t, w.Body.String(), test.expectedResponseBody)
			}
		})
	}
}
