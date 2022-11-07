package group

import (
	"vhosting/pkg/user"

	"github.com/gin-gonic/gin"
)

type Group struct {
	Id   int    `json:"id"   db:"id"`
	Name string `json:"name" db:"name"`
}

type GroupIds struct {
	Ids []int `json:"groupIds" db:"group_id"`
}

type GroupCommon interface {
	CreateGroup(grp *Group) error
	GetGroup(id int) (*Group, error)
	GetAllGroups(urlparams *user.Pagin) (map[int]*Group, error)
	PartiallyUpdateGroup(grp *Group) error
	DeleteGroup(id int) error

	GetUserGroups(id int, urlparams *user.Pagin) (*GroupIds, error)

	IsGroupExists(idOrName interface{}) (bool, error)
}

type GroupUseCase interface {
	GroupCommon

	SetUserGroups(id int, groupIds *GroupIds) error
	DeleteUserGroups(id int, groupIds *GroupIds) error

	BindJSONGroup(ctx *gin.Context) (*Group, error)
	BindJSONGroupIds(ctx *gin.Context) (*GroupIds, error)
	IsGroupRequiredEmpty(name string) bool
	IsGroupIdsRequiredEmpty(groupIds *GroupIds) bool
	AtoiRequestedId(ctx *gin.Context) (int, error)
}

type GroupRepository interface {
	GroupCommon

	SetUserGroups(values string) error
	DeleteUserGroups(id int, condIds string) error
}
