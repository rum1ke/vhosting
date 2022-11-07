package user

import (
	"github.com/gin-gonic/gin"
	"vhosting/pkg/auth"
)

type User struct {
	Id           int    `json:"id"           db:"id"`
	Username     string `json:"username"     db:"username"`
	PasswordHash string `json:"password"     db:"password_hash"`
	IsActive     bool   `json:"isActive"    db:"is_active"`
	IsSuperuser  bool   `json:"isSuperuser" db:"is_superuser"`
	IsStaff      bool   `json:"isStaff"     db:"is_staff"`
	FirstName    string `json:"firstName"   db:"first_name"`
	LastName     string `json:"lastName"    db:"last_name"`
	JoiningDate  string `json:"joiningDate" db:"joining_date"`
	LastLogin    string `json:"lastLogin"   db:"last_login"`
}

type Pagin struct {
	Limit int
	Page  int
}

type UserCommon interface {
	CreateUser(usr *User) error
	GetUser(id int) (*User, error)
	GetAllUsers(urlparams *Pagin) (map[int]*User, error)
	UpdateUserPassword(namepass *auth.Namepass) error
	PartiallyUpdateUser(usr *User) error
	DeleteUser(id int) error

	IsUserSuperuserOrStaff(username string) (bool, error)
	IsUserHavePersonalPermission(userId int, userPerm string) (bool, error)
	IsUserExists(idOrUsername interface{}) (bool, error)
	GetUserId(username string) (int, error)
}

type UserUseCase interface {
	UserCommon

	BindJSONUser(ctx *gin.Context) (*User, error)
	IsRequiredEmpty(username, password string) bool
	AtoiRequestedId(ctx *gin.Context) (int, error)
	ParseURLParams(ctx *gin.Context) *Pagin
}

type UserRepository interface {
	UserCommon
}
