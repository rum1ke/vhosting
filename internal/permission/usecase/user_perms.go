package usecase

import (
	"fmt"

	perm "vhosting/internal/permission"
	"vhosting/pkg/user"
)

func (u *PermUseCase) SetUserPermissions(id int, permIds *perm.PermIds) error {
	values := ""
	for _, val := range permIds.Ids {
		values += fmt.Sprintf("(%d,%d),", id, val)
	}
	values = values[:len(values)-1]
	return u.permRepo.SetUserPermissions(values)
}

func (u *PermUseCase) GetUserPermissions(id int, urlparams *user.Pagin) (*perm.PermIds, error) {
	return u.permRepo.GetUserPermissions(id, urlparams)
}

func (u *PermUseCase) DeleteUserPermissions(id int, permIds *perm.PermIds) error {
	condIds := ""
	for _, val := range permIds.Ids {
		condIds += fmt.Sprintf("%d,", val)
	}
	condIds = condIds[:len(condIds)-1]
	return u.permRepo.DeleteUserPermissions(id, condIds)
}
