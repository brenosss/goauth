package http

import (
	auth "backend/pkg/http/auth"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestAuthenticationMiddleware(t *testing.T) {
	t.Run("UserIsAuthenticated", func(t *testing.T) {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Set("user_id", 1)
		auth.AuthenticationMiddleware()(c)
		assert.Equal(t, 200, w.Result().StatusCode)
	})
	t.Run("UserIsNotAuthenticated", func(t *testing.T) {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Set("user_id", nil)
		auth.AuthenticationMiddleware()(c)
		assert.Equal(t, 401, w.Result().StatusCode)
	})
	t.Run("UserDoesNotExists", func(t *testing.T) {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		auth.AuthenticationMiddleware()(c)
		assert.Equal(t, 401, w.Result().StatusCode)
	})
}
