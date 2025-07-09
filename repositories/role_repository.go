package repositories

import (
	"fmt"
	"candra/backend-api/database"
	"candra/backend-api/models"
)

// Tampilkan semua role
func FindAllRoles(roles *[]models.Role, limit, offset int) error {
	return database.DB.Limit(limit).Offset(offset).Find(roles).Error
}
func FindRoleByID(id string, role *models.Role) error {
	return database.DB.First(role, "id = ?", id).Error
}
func CreateRole(role *models.Role) error {
	return database.DB.Create(role).Error
}
func UpdateRole(id string, updatedData map[string]interface{}) error {
	return database.DB.Model(&models.Role{}).Where("id = ?", id).Updates(updatedData).Error
}
func DeleteRole(id string) error {
	var count int64
	err := database.DB.Model(&models.User{}).Where("role_id = ?", id).Count(&count).Error
	if err != nil {
		return err
	}
	if count > 0 {
		return fmt.Errorf("Role sedang dipakai oleh user, tidak bisa dihapus")
	}
	return database.DB.Delete(&models.Role{}, "id = ?", id).Error
}
