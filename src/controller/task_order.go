package controller

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"todo.app/model"
	"todo.app/service"
)

type taskOrderRequest struct {
	UserID    *int   `json:"userid" binding:"required"`
	OrderList string `json:"orderlist"`
}

// タスクの順序を取得
func TaskOrderGet(userID int) ([]int, error) {

	// DBからデータを取得する
	taskOrderService := service.TaskOrderService{}
	order, err := taskOrderService.GetTaskOrder(userID)
	if err != nil {
		return nil, err
	}

	// ,で区切り配列に入れる
	orderList := strings.Split(order.OrderList, ",")

	// orderをint[]に変換する
	res := make([]int, len(orderList))
	for i := range res {
		res[i], _ = strconv.Atoi(orderList[i])
	}
	return res, nil

}

// DBへタスク順序レコードを追加
// これだけAPIとしては公開しない(関数から呼び出すため)
func TaskOrderAdd(userID int) error {

	// DBへ挿入用にモデルを構築
	taskOrder := model.TaskOrder{
		UserID: userID,
	}
	taskOrderService := service.TaskOrderService{}
	if err := taskOrderService.AddTaskOrder(&taskOrder); err != nil {
		return err
	}

	return nil
}

// タスク順序情報の更新
func TaskOrderUpdate(c *gin.Context) {

	// ポインタを渡してリクエスト内容をバインドする
	taskOrderReq := taskOrderRequest{}
	if err := c.Bind(&taskOrderReq); err != nil {
		c.String(http.StatusBadRequest, "Bad request")
		return
	}

	// DBへ挿入用にモデルを構築
	taskOrder := model.TaskOrder{
		OrderList: taskOrderReq.OrderList,
		// TODO: userid使うときに修正
		UserID: *taskOrderReq.UserID,
	}

	// DBへの書き込み
	taskOrderService := service.TaskOrderService{}
	if err := taskOrderService.UpdateTaskOrder(&taskOrder); err != nil {
		c.String(http.StatusInternalServerError, "Server Error")
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status": "ok",
	})
}
