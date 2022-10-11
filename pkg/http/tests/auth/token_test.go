package http

import (
	"backend/pkg/config"
	db "backend/pkg/database"
	entities "backend/pkg/database/entities"
	auth "backend/pkg/http/auth"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	uuid "github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	config.LoadEnv("test")
	db.ClearDatabase()
	db.ApplyMigrations()
	os.Exit(m.Run())
}

func ErrorContains(out error, want string) bool {
	if out == nil {
		return want == ""
	}
	if want == "" {
		return false
	}
	return strings.Contains(out.Error(), want)
}

func TestValidateToken(t *testing.T) {
	t.Run("ValidToken", func(t *testing.T) {
		content := uuid.NewString()
		token := entities.Token{Content: content, UserId: 1}
		entities.CreateToken(&token)

		got, _ := auth.ValidateToken(content)
		want := content
		assert.Equal(t, want, got.Content)
	})
	t.Run("InvalidToken", func(t *testing.T) {
		_, err := auth.ValidateToken("invalid")
		want := "invalid token"
		if !ErrorContains(err, want) {
			t.Errorf("unexpected error: %v", err)
		}
	})
}

func TestGetRequestToken(t *testing.T) {
	t.Run("HeaderIsValid", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set("Authorization", "Bearer 12345")
		got, err := auth.GetRequestToken(req)
		want := "12345"
		if err != nil {
			t.Errorf("Error is not nil")
		}
		assert.Equal(t, want, got)
	})
	t.Run("HeaderIsInvalid", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set("Authorization", "INVALID 12345")
		_, err := auth.GetRequestToken(req)
		want := "invalid authorization header"
		if !ErrorContains(err, want) {
			t.Errorf("unexpected error: %v", err)
		}
	})
	t.Run("MissingAuthHeader", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		_, err := auth.GetRequestToken(req)
		want := "missing authorization header"
		if !ErrorContains(err, want) {
			t.Errorf("unexpected error: %v", err)
		}
	})
}

func TestTokenAuthMiddleware(t *testing.T) {
	t.Run("AddUserToContext", func(t *testing.T) {
		// DB setup
		tokenContent := uuid.NewString()
		token := entities.Token{Content: tokenContent, UserId: 10}
		entities.CreateToken(&token)

		// Request setup
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set("Authorization", "Bearer "+tokenContent)
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request = req
		auth.TokenAuthMiddleware()(c)

		// Assert
		got, _ := c.Get("user_id")
		assert.Equal(t, 10, got)
	})

	t.Run("InvalidToken", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set("Authorization", "Bearer 9182392183")
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request = req
		auth.TokenAuthMiddleware()(c)
		got, _ := c.Get("user_id")
		assert.Equal(t, nil, got)
	})

	t.Run("WithOutHeader", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request = req
		auth.TokenAuthMiddleware()(c)
		got, _ := c.Get("user_id")
		assert.Equal(t, nil, got)
	})
}
