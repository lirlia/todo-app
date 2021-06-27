package service

import (
	"fmt"

	"todo.app/model"
)

type TaskOrderService struct{}

// TaskOrderの取得
func (TaskOrderService) GetTaskOrder(userID int) (*model.TaskOrder, error) {

	// TODO: 0はuserid
	var taskOrder *model.TaskOrder
	result := db.Select("order_list").Where("user_id = ?", userID).First(&taskOrder)
	if result.Error != nil {
		fmt.Println(result.Error)
		return taskOrder, result.Error
	}
	return taskOrder, nil
}

// TaskOrderの更新
func (TaskOrderService) UpdateTaskOrder(taskOrder *model.TaskOrder) error {

	// TODO: userとcategoryも考慮する
	result := db.Model(&model.TaskOrder{}).Where("user_id = ?", taskOrder.UserID).Updates(taskOrder)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

// TaskOrderの新規追加
// TODO: Userの初回作成時のみ実行される
func (TaskOrderService) AddTaskOrder(taskOrder *model.TaskOrder) error {
	result := db.Create(&taskOrder)

	if result.Error != nil {
		return result.Error
	}
	return nil
}
