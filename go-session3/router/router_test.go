package router_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"training-go/go-session3/router"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestSetupRouter_RootHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	router.SetupRouter(r)

	//request GET ke /
	req, _ := http.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	//cek status code
	assert.Equal(t, http.StatusOK, w.Code)

	expectedBody := `{"message":"Hello, dari Gin"}`
	assert.JSONEq(t, expectedBody, w.Body.String())

}
func TestSetupRouter_PositiveCase(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	router.SetupRouter(r)

	//persiapkan data json
	requestbody := map[string]string{"message": "test message"}
	requestbodyBytes, _ := json.Marshal(requestbody)

	//request GET ke /private dengan data json valid
	req, _ := http.NewRequest("POST", "/private/post", bytes.NewBuffer(requestbodyBytes))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "valid-token")

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	//cek status code
	assert.Equal(t, http.StatusOK, w.Code)

	expectedBody := `{"message":"test message"}`
	assert.JSONEq(t, expectedBody, w.Body.String())
}

func TestSetupRouter_NegativeCase_BadRequest(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	router.SetupRouter(r)

	//request GET ke /private dengan data json yg tidak valid
	req, _ := http.NewRequest("POST", "/private/post", bytes.NewBufferString("{Invalid JSON}"))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "valid-token")

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	//cek status code
	assert.Contains(t, w.Body.String(), "invalid character")
}

func TestSetupRouter_NegativeCase_Unauthorized(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	router.SetupRouter(r)

	//request GET ke /private dengan token tidak valid
	req, _ := http.NewRequest("POST", "/private/post", nil)
	//req.Header.Set("Authorization", "invalid-token")

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	//cek status code
	assert.Equal(t, http.StatusUnauthorized, w.Code)
	assert.Contains(t, w.Body.String(), `{"error":"Authorization token required"}`)
}
