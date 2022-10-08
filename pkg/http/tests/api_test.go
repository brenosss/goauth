package http

import (
	"backend/pkg/config"
	http "backend/pkg/http"
	netHttp "net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	config.LoadEnv("test")
	os.Exit(m.Run())
}

func TestAuthenticationMiddleware(t *testing.T) {
	t.Run("UserIsAuthenticated", func(t *testing.T) {
		router := http.SetupRouter()
		w := httptest.NewRecorder()
		req, _ := netHttp.NewRequest("GET", "/ping", nil)
		req.Header.Set("Authorization", "Bearer 12345")
		router.ServeHTTP(w, req)

		assert.Equal(t, 200, w.Code)
		assert.Equal(t, "1 pong\n", w.Body.String())
	})
	t.Run("UserIsUnauthenticated", func(t *testing.T) {
		router := http.SetupRouter()
		w := httptest.NewRecorder()
		req, _ := netHttp.NewRequest("GET", "/ping", nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, 401, w.Code)
		assert.Equal(t, "Unauthorized", w.Body.String())
	})
}
