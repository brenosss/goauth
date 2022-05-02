package http

import (
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type Token struct {
	Id         int
	Token      string
	Created_at time.Time
	User_id    int
}

func ValidateToken(token string) (*Token, error) {
	// This needs to come from DB
	validToken := Token{Id: 1, Token: "12345", Created_at: time.Now(), User_id: 1}
	if token == validToken.Token {
		return &validToken, nil
	}
	return nil, errors.New("invalid token")
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
				c.Set("user_id", token.User_id)
			}
		}
		c.Next()
	}
}
