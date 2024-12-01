package impl

import (
	"maulanadityaa/laundry-app-rest-api/config"
	"maulanadityaa/laundry-app-rest-api/model/entity"
)

type RoleRepository struct{}

func NewRoleRepository() *RoleRepository {
	return &RoleRepository{}
}

func (RoleRepository) GetRoleByID(roleID string) (entity.Role, error) {
	var role entity.Role

	if err := config.DB.Where("id = ?", roleID).First(&role).Error; err != nil {
		return entity.Role{}, err
	}

	return role, nil
}

func (RoleRepository) GetRoleByName(roleName string) (entity.Role, error) {
	var role entity.Role

	if err := config.DB.Where("name = ?", roleName).First(&role).Error; err != nil {
		return entity.Role{}, err
	}

	return role, nil
}
