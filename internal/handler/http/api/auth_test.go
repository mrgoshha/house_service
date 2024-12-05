package api

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	httpApi "houseService/internal/handler/http"
	mockmanager "houseService/pkg/auth/mocks"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestAuthManager_userIdentity(t *testing.T) {
	// Init Test Table
	type mockBehavior func(r *mockmanager.MockTokenManager, token string)

	testTable := []struct {
		name                 string
		headerName           string
		headerValue          string
		token                string
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:        "Ok",
			headerName:  "Authorization",
			headerValue: "Bearer token",
			token:       "token",
			mockBehavior: func(r *mockmanager.MockTokenManager, token string) {
				r.EXPECT().Parse(token).Return("1", "user", nil)
			},
			expectedStatusCode: 200,
			expectedResponseBody: `{"user_type":"user"}
`,
		},
		{
			name:                 "Invalid Header Name",
			headerName:           "",
			headerValue:          "Bearer token",
			token:                "token",
			mockBehavior:         func(r *mockmanager.MockTokenManager, token string) {},
			expectedStatusCode:   401,
			expectedResponseBody: `"message":"empty auth header"`,
		},
		{
			name:                 "Invalid Header Value",
			headerName:           "Authorization",
			headerValue:          "B token",
			token:                "token",
			mockBehavior:         func(r *mockmanager.MockTokenManager, token string) {},
			expectedStatusCode:   401,
			expectedResponseBody: `"message":"invalid auth header"`,
		},
		{
			name:                 "Empty Token",
			headerName:           "Authorization",
			headerValue:          "Bearer ",
			token:                "token",
			mockBehavior:         func(r *mockmanager.MockTokenManager, token string) {},
			expectedStatusCode:   401,
			expectedResponseBody: `"message":"token is empty"`,
		},
		{
			name:        "Parse Error",
			headerName:  "Authorization",
			headerValue: "Bearer token",
			token:       "token",
			mockBehavior: func(r *mockmanager.MockTokenManager, token string) {
				r.EXPECT().Parse(token).Return("", "", errors.New("invalid token"))
			},
			expectedStatusCode:   401,
			expectedResponseBody: `"message":"invalid token"`,
		},
	}

	for _, test := range testTable {
		t.Run(test.name, func(t *testing.T) {
			// Init Dependencies
			c := gomock.NewController(t)
			defer c.Finish()

			manager := mockmanager.NewMockTokenManager(c)
			test.mockBehavior(manager, test.token)

			authManager := NewTokenManager(manager)

			log := slog.New(
				slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))

			// Init Endpoint
			r := httpApi.NewRouter(log)
			r.Use(authManager.UserIdentity)
			r.HandleFunc("/identity", func(w http.ResponseWriter, r *http.Request) {
				data := struct {
					UserType string `json:"user_type"`
				}{
					UserType: r.Context().Value(CtxKeyUserType).(string),
				}
				response(w, r, http.StatusOK, data)
			}).Methods(http.MethodGet)

			// Init Test Request
			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", "/identity", nil)
			req.Header.Set(test.headerName, test.headerValue)

			r.ServeHTTP(w, req)

			// Asserts
			assert.Equal(t, w.Code, test.expectedStatusCode)
			if w.Code == http.StatusOK {
				assert.Equal(t, w.Body.String(), test.expectedResponseBody)
			} else {
				assert.Contains(t, w.Body.String(), test.expectedResponseBody)
			}
		})
	}
}
