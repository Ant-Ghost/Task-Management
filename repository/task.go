package repository

import (
	"errors"
	"zocket/task_manager/database"
	"zocket/task_manager/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func CreateTask(task *models.Task) (*models.Task, error) {
	if err := database.DB.Create(task).Error; err != nil {
		return nil, err
	}
	return task, nil
}
func GetTaskByID(id uint) (*models.Task, error) {
	var task models.Task
	if err := database.DB.First(&task, id).Error; err != nil {
		return nil, err
	}
	return &task, nil
}

func GetAllTasks() ([]models.Task, error) {
	var tasks []models.Task
	if err := database.DB.Find(&tasks).Error; err != nil {
		return nil, err
	}
	return tasks, nil
}
func UpdateTask(task *models.Task) (*models.Task, error) {

	tx := database.DB.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}
	defer tx.Rollback() // Rollback the transaction if not committed

	// Attempt to acquire a row-level lock using SELECT FOR UPDATE
	var existingTask models.Task
	if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).First(&existingTask, task.ID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("task not found")
		}
		return nil, err
	}

	if err := tx.Save(task).Error; err != nil {
		return nil, err
	}
	return task, tx.Commit().Error
}

func UpdateTaskFields(taskId uint, fields map[string]interface{}) (*models.Task, error) {
	var task models.Task
	if err := database.DB.Model(&task).Where("id = ?", taskId).Updates(fields).Error; err != nil {
		return nil, err
	}
	return &task, nil
}
func DeleteTask(id uint) error {
	if err := database.DB.Delete(&models.Task{}, id).Error; err != nil {
		return err
	}
	return nil
}

func GetTasksByFilters(
	filters map[string]interface{},
	searchQuery string,
	page, pageSize int,
) ([]models.Task, error) {
	var tasks []models.Task

	query := database.DB.Model(&models.Task{})

	for key, value := range filters {
		query = query.Where(key+" = ?", value)
	}
	if searchQuery != "" {
		query = query.Where("title LIKE ? OR description LIKE ?", "%"+searchQuery+"%", "%"+searchQuery+"%")
	}
	query = query.Offset((page - 1) * pageSize).Limit(pageSize)
	if err := query.Find(&tasks).Error; err != nil {
		return nil, err
	}
	return tasks, nil
}

func GetTasksByParentID(parentID uint) ([]models.Task, error) {
	var tasks []models.Task
	if err := database.DB.Where("parent_id = ?", parentID).Find(&tasks).Error; err != nil {
		return nil, err
	}
	return tasks, nil
}

func GetTasksByStatus(status string) ([]models.Task, error) {
	var tasks []models.Task
	if err := database.DB.Where("status = ?", status).Find(&tasks).Error; err != nil {
		return nil, err
	}
	return tasks, nil
}
