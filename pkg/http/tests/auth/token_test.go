package http

import (
	auth "backend/pkg/http/auth"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

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
		got, _ := auth.ValidateToken("12345")
		want := "12345"
		assert.Equal(t, want, got.Token)
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
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set("Authorization", "Bearer 12345")
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request = req
		auth.TokenAuthMiddleware()(c)
		got, _ := c.Get("user_id")
		assert.Equal(t, 1, got)
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
