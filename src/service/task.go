package service

import (
	"todo.app/model"
)

type TaskService struct{}

// TaskのInsert
func (TaskService) SetTask(Task *model.Task) error {
	result := db.Create(&Task)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// Taskの全取得
func (TaskService) GetTaskList() ([]model.Task, error) {

	rows, err := db.Model(&model.Task{}).Where("user_id = ?", 0).Rows()
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var task model.Task
	var tasks []model.Task

	for rows.Next() {
		// ScanRowsは1行をtaskに変換します
		db.ScanRows(rows, &task)
		tasks = append(tasks, task)
	}
	return tasks, nil
}

func (TaskService) UpdateTask(newTask *model.Task) error {
	// TODO: 後で実装
	// https://gorm.io/ja_JP/docs/update.html
	return nil
}

func (TaskService) DeleteTask(id int) error {
	// TODO: 後で実装
	// https://gorm.io/ja_JP/docs/delete.html
	return nil
}
