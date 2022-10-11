package http

import (
	entities "backend/pkg/database/entities"
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func ValidateToken(token string) (*entities.Token, error) {
	// This needs to come from DB
	validToken := entities.GetToken(token)
	if validToken.Id == 0 {
		return nil, errors.New("invalid token")
	}
	return &validToken, nil
}

func GetRequestToken(r *http.Request) (string, error) {
	authorizationHeader := r.Header.Get("Authorization")
	if authorizationHeader == "" {
		return "", errors.New("missing authorization header")
	}
	headerValue := strings.Split(authorizationHeader, "Bearer ")
	if len(headerValue) != 2 {
		return "", errors.New("invalid authorization header")
	}
	token := headerValue[1]
	return token, nil
}

func TokenAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestToken, err := GetRequestToken(c.Request)
		if err != nil {
			c.Set("user_id", nil)
		} else {
			token, err := ValidateToken(requestToken)
			if err != nil {
				c.Set("user_id", nil)
			} else {
				c.Set("user_id", token.UserId)
			}
		}
		c.Next()
	}
}
