package model

import (
	"time"

	"todo.app/conf"
)

type Task struct {
	TaskID    int `gorm:"primaryKey"`
	Name      string
	Done      bool
	Message   string
	UserID    int `gorm:"foreignKey"`
	CreatedAt time.Time
}

func (Task) TableName() string {
	return conf.DB_TASK_TABLE_NAME
}
