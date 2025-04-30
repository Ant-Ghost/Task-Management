package repository

import (
	"zocket/task_manager/database"
	"zocket/task_manager/models"
)

func CreateUser(user *models.User) (*models.User, error) {

	if err := database.DB.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func GetUserByID(id uint) (*models.User, error) {
	var user models.User
	if err := database.DB.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func GetUserByUsername(username string) (*models.User, error) {
	var user models.User
	if err := database.DB.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func GetAllUsers() ([]models.User, error) {
	var users []models.User
	if err := database.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func UpdateUser(user *models.User) (*models.User, error) {
	if err := database.DB.Save(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}
func DeleteUser(id uint) error {
	if err := database.DB.Delete(&models.User{}, id).Error; err != nil {
		return err
	}
	return nil
}

// Search Task by Title or Description
func SearchTasks(query string) ([]models.Task, error) {
	var tasks []models.Task
	if err := database.DB.Where("title LIKE ? OR description LIKE ?", "%"+query+"%", "%"+query+"%").Find(&tasks).Error; err != nil {
		return nil, err
	}
	return tasks, nil
}
