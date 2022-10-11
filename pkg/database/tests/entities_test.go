package database

import (
	"backend/pkg/config"
	db "backend/pkg/database"
	"backend/pkg/database/entities"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	config.LoadEnv("test")
	db.ClearDatabase()
	db.ApplyMigrations()
	os.Exit(m.Run())
}

func TestToken(t *testing.T) {
	t.Run("Create a new token", func(t *testing.T) {
		token := entities.Token{Content: "12345", UserId: 1}
		entities.CreateToken(&token)

		tokenDb := entities.GetToken(token.Content)
		assert.Equal(t, token.Content, tokenDb.Content)
		assert.Equal(t, token.UserId, tokenDb.UserId)
	})
	t.Run("Token does not exists return empty token", func(t *testing.T) {
		token := entities.Token{Content: "12345", UserId: 1}
		entities.CreateToken(&token)

		tokenDb := entities.GetToken("RandomContent")
		assert.Equal(t, entities.Token{}, tokenDb)
	})
}
