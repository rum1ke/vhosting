package info

import (
	"github.com/gin-gonic/gin"
	"vhosting/pkg/user"
)

type Info struct {
	Id          int    `json:"id"           db:"id"`
	CreateDate  string `json:"createDate"   db:"create_date"`
	Stream      string `json:"stream"       db:"stream"`
	StartPeriod string `json:"startPeriod"  db:"start_period"`
	StopPeriod  string `json:"stopPeriod"   db:"stop_period"`
	TimeLife    string `json:"timeLife"     db:"time_life"`
	UserId      int    `json:"userId"       db:"user_id"`
}

type InfoCommon interface {
	CreateInfo(nfo *Info) error
	GetInfo(id int) (*Info, error)
	GetAllInfos(urlparams *user.Pagin) (map[int]*Info, error)
	PartiallyUpdateInfo(nfo *Info) error
	DeleteInfo(id int) error

	IsInfoExists(id int) (bool, error)
}

type InfoUseCase interface {
	InfoCommon

	BindJSONInfo(ctx *gin.Context) (*Info, error)
	IsRequiredEmpty(stream string) bool
	AtoiRequestedId(ctx *gin.Context) (int, error)
}

type InfoRepository interface {
	InfoCommon
}
