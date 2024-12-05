package tests

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
)

func (s *APITestSuite) TestStudentCreateFlat() {
	r := s.Require()

	//Arrange
	houseId, price, rooms := 1, 34, 3
	inputFlat := fmt.Sprintf(`{"house_id":%d,"price":%d,"rooms":%d}`, houseId, price, rooms)
	outputFlat := fmt.Sprintf(`"house_id":%d,"price":%d,"rooms":%d,"status":"created"}`, houseId, price, rooms)

	t, _ := s.tokenManager.NewJWT("1", "client")
	token := fmt.Sprintf("Bearer %s", t)

	// Create Request
	req := httptest.NewRequest("POST", "/flat/create",
		bytes.NewBufferString(inputFlat))
	req.Header.Set("Content-type", "application/json")
	req.Header.Set("Authorization", token)

	// Make Request
	resp := httptest.NewRecorder()
	s.serviceProvider.HttpRouter().ServeHTTP(resp, req)

	//Assert
	r.Equal(http.StatusOK, resp.Result().StatusCode)
	r.Contains(resp.Body.String(), outputFlat)

}

func (s *APITestSuite) TestStudentCreateFlatNotAuth() {
	r := s.Require()

	//Arrange
	houseId, price, rooms := 1, 34, 3
	inputFlat := fmt.Sprintf(`{"house_id":%d,"price":%d,"rooms":%d}`, houseId, price, rooms)

	// Create Request
	req := httptest.NewRequest("POST", "/flat/create",
		bytes.NewBufferString(inputFlat))
	req.Header.Set("Content-type", "application/json")

	// Make Request
	resp := httptest.NewRecorder()
	s.serviceProvider.HttpRouter().ServeHTTP(resp, req)

	//Assert
	r.Equal(http.StatusUnauthorized, resp.Result().StatusCode)

}

func (s *APITestSuite) TestHouseIdGetClient() {
	r := s.Require()

	//Arrange
	houseId, price, rooms := 1, 30, 2
	outputFlat := fmt.Sprintf(`"house_id":%d,"price":%d,"rooms":%d,"status":"approved"}`, houseId, price, rooms)

	t, _ := s.tokenManager.NewJWT("1", "client")
	token := fmt.Sprintf("Bearer %s", t)

	// Create Request
	req, _ := http.NewRequest(http.MethodGet, "/house?id=1", nil)
	req.Header.Set("Content-type", "application/json")
	req.Header.Set("Authorization", token)

	// Make Request
	resp := httptest.NewRecorder()
	s.serviceProvider.HttpRouter().ServeHTTP(resp, req)

	//Assert
	r.Equal(http.StatusOK, resp.Result().StatusCode)
	r.Contains(resp.Body.String(), outputFlat)

}

func (s *APITestSuite) TestHouseIdGetModerator() {
	r := s.Require()

	//Arrange
	houseId, price, rooms := 1, 20, 1
	outputFlat := fmt.Sprintf(`"house_id":%d,"price":%d,"rooms":%d,"status":"created"}`, houseId, price, rooms)

	t, _ := s.tokenManager.NewJWT("1", "moderator")
	token := fmt.Sprintf("Bearer %s", t)

	// Create Request
	req, _ := http.NewRequest(http.MethodGet, "/house?id=1", nil)
	req.Header.Set("Content-type", "application/json")
	req.Header.Set("Authorization", token)

	// Make Request
	resp := httptest.NewRecorder()
	s.serviceProvider.HttpRouter().ServeHTTP(resp, req)

	//Assert
	r.Equal(http.StatusOK, resp.Result().StatusCode)
	r.Contains(resp.Body.String(), outputFlat)

}

func (s *APITestSuite) TestHouseIdGetNotAuth() {
	r := s.Require()

	// Create Request
	req, _ := http.NewRequest(http.MethodGet, "/house?id=1", nil)
	req.Header.Set("Content-type", "application/json")

	// Make Request
	resp := httptest.NewRecorder()
	s.serviceProvider.HttpRouter().ServeHTTP(resp, req)

	//Assert
	r.Equal(http.StatusUnauthorized, resp.Result().StatusCode)
}
