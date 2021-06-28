package model

import (
	"time"

	"todo.app/conf"
)

type Task struct {
	TaskID    int `gorm:"primaryKey"`
	Title     string
	Done      *bool `gorm:"default:false"`
	Message   string
	UserID    *int `gorm:"foreignKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (Task) TableName() string {
	return conf.DB_TASK_TABLE_NAME
}

type TaskOrder struct {
	UserID    int `gorm:"primaryKey,foreignKey"`
	OrderList string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (TaskOrder) TableName() string {
	return conf.DB_TASKORDER_TABLE_NAME
}
