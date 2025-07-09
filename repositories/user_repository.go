package repositories

import (
	"candra/backend-api/database"
	"candra/backend-api/models"
)

func GetAllUsers(users *[]models.User, limit, offset int) error {
	return database.DB.Preload("Role").Limit(limit).Offset(offset).Find(users).Error
}

func CreateUser(user *models.User) error {
	return database.DB.Create(user).Error
}

func FindUserByUsername(username string, user *models.User) error {
	return database.DB.Where("username = ?", username).First(user).Error
}

// FindUserByID
func FindUserByID(id string, user *models.User) error {
	return database.DB.Preload("Role").First(user, "id = ?", id).Error
}

func CountAllUsers() (int64, error) {
	var count int64
	err := database.DB.Model(&models.User{}).Count(&count).Error
	return count, err
}
func UpdateUser(user *models.User) error {
	return database.DB.Save(user).Error
}
func DeleteUser(user *models.User) error {
	return database.DB.Delete(user).Error
}
