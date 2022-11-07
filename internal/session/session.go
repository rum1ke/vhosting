package session

import "github.com/gin-gonic/gin"

type Session struct {
	Id           int    `db:"id"`
	Content      string `db:"content"`
	CreationDate string `db:"creationDate"`
}

type SessCommon interface {
	GetSessionAndDate(token string) (*Session, error)
	DeleteSession(token string) error
}

type SessUseCase interface {
	SessCommon

	CreateSession(ctx *gin.Context, username, token, timestamp string) error
}

type SessRepository interface {
	SessCommon

	CreateSession(session *Session) error
}
