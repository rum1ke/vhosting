package permission

import (
	"vhosting/pkg/user"

	"github.com/gin-gonic/gin"
)

type Perm struct {
	Id       int    `json:"id"        db:"id"`
	Name     string `json:"name"      db:"name"`
	Codename string `json:"codeName" db:"code_name"`
}

type PermIds struct {
	Ids []int `json:"permIds" db:"perm_id"`
}

type PermCommon interface {
	GetAllPermissions(urlparams *user.Pagin) (map[int]*Perm, error)
	GetUserPermissions(id int, urlparams *user.Pagin) (*PermIds, error)
	GetGroupPermissions(id int, urlparams *user.Pagin) (*PermIds, error)
}

type PermUseCase interface {
	PermCommon

	SetUserPermissions(id int, permIds *PermIds) error
	DeleteUserPermissions(id int, permIds *PermIds) error

	SetGroupPermissions(id int, permIds *PermIds) error
	DeleteGroupPermissions(id int, permIds *PermIds) error

	BindJSONPermIds(ctx *gin.Context) (*PermIds, error)
	IsRequiredEmpty(permIds *PermIds) bool
	AtoiRequestedId(ctx *gin.Context) (int, error)
}

type PermRepository interface {
	PermCommon

	SetUserPermissions(values string) error
	DeleteUserPermissions(id int, condIds string) error

	SetGroupPermissions(values string) error
	DeleteGroupPermissions(id int, condIds string) error
}
