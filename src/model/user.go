package model

import (
	"time"

	"todo.app/conf"
)

type User struct {
	UserID    int `gorm:"primaryKey"`
	Name      string
	Password  string
	CreatedAt time.Time
}

// テーブル名を決定する
func (User) TableName() string {
	return conf.DB_USER_TABLE_NAME
}
