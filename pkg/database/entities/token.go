package entities

import (
	db "backend/pkg/database"
	"time"

	types "database/sql"
)

type Token struct {
	Id      int
	Content string
	UserId  int

	CreatedAt time.Time
	UpdatedAt types.NullTime
}

func CreateToken(token *Token) string {
	query := db.GetQuery("token/create_token.sql")
	_, err := db.GetConnection().Exec(query, token.Content, token.UserId)
	if err != nil {
		panic(err)
	}
	return token.Content
}

func GetToken(value string) Token {
	query := db.GetQuery("token/get_token.sql")
	var token Token
	var err = db.GetConnection().QueryRow(query, value).Scan(&token.Id, &token.Content, &token.UserId, &token.CreatedAt, &token.UpdatedAt)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return token
		}
		panic(err)
	}
	return token
}
