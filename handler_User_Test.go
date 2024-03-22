package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Calmantara/go-kominfo-2024/blob/main/go-middleware/internal/handler/user.go"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestDeleteUsersById(t *testing.T) {
	t.Run("invalid required param", func(t *testing.T) {
		// Setup
		gin.SetMode(gin.TestMode)
		router := gin.New()
		router.DELETE("/users/:id", DeleteUsersById)

		req, _ := http.NewRequest("DELETE", "/users/invalid", nil)
		w := httptest.NewRecorder()

		// Execute
		router.ServeHTTP(w, req)

		// Verify
		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("invalid user session", func(t *testing.T) {
		// Setup
		gin.SetMode(gin.TestMode)
		router := gin.New()
		router.DELETE("/users/:id", DeleteUsersById)

		req, _ := http.NewRequest("DELETE", "/users/123", nil)
		w := httptest.NewRecorder()

		// Execute
		router.ServeHTTP(w, req)

		// Verify
		assert.Equal(t, http.StatusUnauthorized, w.Code)
	})

	t.Run("invalid user id session", func(t *testing.T) {
		// Setup
		gin.SetMode(gin.TestMode)
		router := gin.New()
		router.DELETE("/users/:id", DeleteUsersById)

		req, _ := http.NewRequest("DELETE", "/users/123", nil)
		req.Header.Set("User-Id", "invalid")
		w := httptest.NewRecorder()

		// Execute
		router.ServeHTTP(w, req)

		// Verify
		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("invalid user request", func(t *testing.T) {
		// Setup
		gin.SetMode(gin.TestMode)
		router := gin.New()
		router.DELETE("/users/:id", DeleteUsersById)

		req, _ := http.NewRequest("DELETE", "/users/123", nil)
		req.Header.Set("User-Id", "456")
		w := httptest.NewRecorder()

		// Execute
		router.ServeHTTP(w, req)

		// Verify
		assert.Equal(t, http.StatusUnauthorized, w.Code)
	})

	t.Run("user not found", func(t *testing.T) {
		// Setup
		gin.SetMode(gin.TestMode)
		router := gin.New()
		router.DELETE("/users/:id", DeleteUsersById)

		req, _ := http.NewRequest("DELETE", "/users/123", nil)
		req.Header.Set("User-Id", "123")
		w := httptest.NewRecorder()

		// Execute
		router.ServeHTTP(w, req)

		// Verify
		assert.Equal(t, http.StatusNotFound, w.Code)
	})

	t.Run("success", func(t *testing.T) {
		// Setup
		gin.SetMode(gin.TestMode)
		router := gin.New()
		router.DELETE("/users/:id", DeleteUsersById)

		req, _ := http.NewRequest("DELETE", "/users/123", nil)
		req.Header.Set("User-Id", "123")
		w := httptest.NewRecorder()

		// Execute
		router.ServeHTTP(w, req)

		// Verify
		assert.Equal(t, http.StatusOK, w.Code)
	})
}
