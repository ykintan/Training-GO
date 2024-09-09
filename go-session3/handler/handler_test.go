package handler_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"training-go/go-session3/handler"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetHelloMessage(t *testing.T) {
	t.Run("Positive Test - correct messgae", func(t *testing.T) {
		expectedOutput := "Hello, dari Gin"
		actualOutput := handler.GetHelloMessage()
		require.Equal(t, expectedOutput, actualOutput, "The two words should be the same.", expectedOutput)
	})
}

func TestRootHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)

	router := gin.Default()
	router.GET("/", handler.RootHandler)

	req, _ := http.NewRequest("GET", "/", nil)

	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	expectedBody := `{"message":"Hello, dari Gin"}`
	assert.JSONEq(t, expectedBody, w.Body.String())
}

type JsonRequest struct {
	Message string `json:"message"`
}

func TestPostHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)
	//setup router
	router := gin.Default()
	router.POST("/", handler.PostHandler)

	t.Run("Positive Case", func(t *testing.T) {
		//persiapkan data json
		requestbody := JsonRequest{Message: "Hello, dari Test!"}
		requestbodyBytes, _ := json.Marshal(requestbody)

		//buat request http post
		req, _ := http.NewRequest("POST", "/", bytes.NewBuffer(requestbodyBytes))
		req.Header.Set("Content-Type", "application/json")

		//buat response recorder untuk merekam response
		w := httptest.NewRecorder()

		//lakukan permintaan
		router.ServeHTTP(w, req)

		//periksa status code
		assert.Equal(t, http.StatusOK, w.Code)

		//periksa body response
		expectedBody := `{"message":"Hello, dari Test!"}`
		assert.JSONEq(t, expectedBody, w.Body.String())
	})

	t.Run("Negative Case", func(t *testing.T) {
		//persiapkan data json
		requestbody := ""
		requestbodyBytes := []byte(requestbody)

		//buat request http post
		req, _ := http.NewRequest("POST", "/", bytes.NewBuffer(requestbodyBytes))
		req.Header.Set("Content-Type", "application/json")

		//buat response recorder untuk merekam response
		w := httptest.NewRecorder()

		//lakukan permintaan
		router.ServeHTTP(w, req)

		//periksa status code
		assert.Equal(t, http.StatusBadRequest, w.Code)

		//periksa body response
		//assert.Contains(t, w.Body.String(), "EOF")
		assert.Contains(t, w.Body.String(), "{\"error\":\"EOF\"}")
	})
}
