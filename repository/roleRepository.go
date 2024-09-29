package repository

import "github.com/maulanadityaa/laundry-app-rest-api/model/entity"

type RoleRepository interface {
	GetRoleByID(roleID string) (entity.Role, error)
	GetRoleByName(roleName string) (entity.Role, error)
}
