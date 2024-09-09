package middleware_test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"training-go/go-session3/middleware"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestAuthMiddleware_PositiveCase(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.Use(middleware.AuthMiddleware())

	//handler yang hanya dapat diakses oleh token valid
	r.GET("/private", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "private data"})
	})

	//request GET ke /private dengan token valid
	req, _ := http.NewRequest(http.MethodGet, "/private", nil)
	req.Header.Set("Authorization", "valid-token")

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	//cek status code
	assert.Equal(t, http.StatusOK, w.Code)
	assert.JSONEq(t, `{"message":"private data"}`, w.Body.String())
}

func TestAuthMiddleware_NegativeCase_Notoken(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.Use(middleware.AuthMiddleware())

	//handler yang hanya dapat diakses oleh token valid
	r.GET("/private", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "private data"})
	})

	//request GET ke /private dengan token valid
	req, _ := http.NewRequest(http.MethodGet, "/private", nil)

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	//cek status code
	assert.Equal(t, http.StatusUnauthorized, w.Code)
	//assert.JSONEq(t, w.Body.String(), `{"error":"Authorization token required"}`)
	assert.Contains(t, w.Body.String(), `{"error":"Authorization token required"}`)
}

func TestAuthMiddleware_NegativeCase_InvalidToken(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.Use(middleware.AuthMiddleware())

	//handler yang hanya dapat diakses oleh token valid
	r.GET("/private", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "private data"})
	})

	//request GET ke /private dengan token valid
	req, _ := http.NewRequest(http.MethodGet, "/private", nil)
	req.Header.Set("Authorization", "invalid-token")

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	//cek status code
	assert.Equal(t, http.StatusUnauthorized, w.Code)
	//assert.JSONEq(t, w.Body.String(), `{"error":"Invalid authorization token"}`)
	assert.Contains(t, w.Body.String(), `{"error":"Invalid authorization token"}`)
}
