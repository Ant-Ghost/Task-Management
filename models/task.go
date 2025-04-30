package models

import (
	"gorm.io/gorm"
)

type TaskStatus int

const (
	Pending TaskStatus = iota
	InProgress
	OnHold
	Completed
	Cancelled
)

func (status TaskStatus) String() string {
	switch status {
	case Pending:
		return "Pending"
	case InProgress:
		return "In Progress"
	case OnHold:
		return "On Hold"
	case Completed:
		return "Completed"
	case Cancelled:
		return "Cancelled"
	default:
		return "Unknown"
	}
}

type TaskPriority int

const (
	Low TaskPriority = iota
	Medium
	High
)

func (priority TaskPriority) String() string {
	switch priority {
	case Low:
		return "Low"
	case Medium:
		return "Medium"
	case High:
		return "High"
	default:
		return "Unknown"
	}
}

// Task represents a task in the task manager.
type Task struct {
	gorm.Model
	Title        string `json:"title" gorm:"not null"`
	Description  string `json:"description"`
	Status       string `json:"status" gorm:"default:'Pending'"`
	ParentID     *uint  `json:"parent_id"`
	Parent       *Task  `json:"parent" gorm:"foreignKey:ParentID"`
	SubTaskCount int    `json:"sub_task_count" gorm:"default:0"`

	AssignToID *uint `json:"assign_to_id"`
	AssignTo   *User `json:"assign_to" gorm:"foreignKey:AssignToID"`

	CreatedByID *uint `json:"created_by_id"`
	CreatedBy   *User `json:"created_by" gorm:"foreignKey:CreatedByID"`

	ReportToID *uint `json:"report_to_id"`
	ReportTo   *User `json:"report_to" gorm:"foreignKey:ReportToID"`

	Priority string `json:"priority" gorm:"default:'Low'"`
}

func (t *Task) AfterCreate(tx *gorm.DB) error {
	if t.ParentID != nil {
		var count int64
		if err := tx.Model(&Task{}).Where("parent_id = ?", t.ParentID).Count(&count).Error; err != nil {
			return err
		}
		if err := tx.Model(&Task{}).Where("id = ?", *t.ParentID).Update("sub_task_count", count).Error; err != nil {
			return err
		}
	}
	return nil
}

func (t *Task) AfterDelete(tx *gorm.DB) error {
	if t.ParentID != nil {
		var count int64
		if err := tx.Model(&Task{}).Where("parent_id = ?", t.ParentID).Count(&count).Error; err != nil {
			return err
		}
		if err := tx.Model(&Task{}).Where("id = ?", *t.ParentID).Update("sub_task_count", count).Error; err != nil {
			return err
		}
	}
	return nil
}
