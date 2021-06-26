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

	rows, err := db.Model(&model.Task{}).Where("user_id = ?", 0).Order("created_at desc").Rows()
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

// Taskのアップデート
// https://gorm.io/ja_JP/docs/update.html
func (TaskService) UpdateTask(Task *model.Task) error {
	// TODO: userも考慮する
	result := db.Model(&model.Task{}).Where("task_id = ?", Task.TaskID).Updates(Task)

	if result.Error != nil {
		return result.Error
	}
	return nil
}

// Taskの削除
// https://gorm.io/ja_JP/docs/delete.html
func (TaskService) DeleteTask(Task *model.Task) error {

	//削除処理を実行
	if err := db.Delete(model.Task{}, Task.TaskID).Error; err != nil {
		return err
	}

	return nil
}
