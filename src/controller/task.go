package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"todo.app/model"
	"todo.app/service"
)

type TaskRequest struct {
	Title string `json:"title" binding:"required,max=100"`
	// https://github.com/gin-gonic/gin/issues/814
	Done    *bool  `json:"done" binding:"required"`
	Message string `json:"message" binding:"max=1000"`
	// https://github.com/gin-gonic/gin/issues/737
	UserID *int `json:"userid" binding:"required"`
}

// DBへタスクを追加
func TaskAdd(c *gin.Context) {

	taskReq := TaskRequest{}

	// ポインタを渡してリクエスト内容をバインドする
	if err := c.Bind(&taskReq); err != nil {
		c.String(http.StatusBadRequest, "Bad request")
		return
	}

	// DBへ挿入用にモデルを構築
	task := model.Task{
		Title:   taskReq.Title,
		Done:    taskReq.Done,
		Message: taskReq.Message,
		// TODO: userid使うときに修正
		UserID: taskReq.UserID,
	}
	TaskService := service.TaskService{}
	err := TaskService.SetTask(&task)
	if err != nil {
		c.String(http.StatusInternalServerError, "Server Error")
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status": "ok",
	})
}

// 登録されているタスクの一覧を取得
func TaskList(c *gin.Context) {
	TaskService := service.TaskService{}
	TaskLists, err := TaskService.GetTaskList()
	if err != nil {
		c.String(http.StatusInternalServerError, "Server Error")
		return
	}
	c.JSONP(http.StatusOK, gin.H{
		"message": "ok",
		"data":    TaskLists,
	})
}

// タスク情報の更新
func TaskUpdate(c *gin.Context) {

	// ポインタを渡してリクエスト内容をバインドする
	taskReq := TaskRequest{}
	if err := c.Bind(&taskReq); err != nil {
		c.String(http.StatusBadRequest, "Bad request")
		return
	}

	// taskidをURLから抽出
	taskid, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.String(http.StatusBadRequest, "Bad request")
		return
	}

	// DBへ挿入用にモデルを構築
	task := model.Task{
		TaskID:  taskid,
		Title:   taskReq.Title,
		Done:    taskReq.Done,
		Message: taskReq.Message,
		// TODO: userid使うときに修正
		UserID: taskReq.UserID,
	}

	// DBへの書き込み
	TaskService := service.TaskService{}
	err = TaskService.UpdateTask(&task)
	if err != nil {
		c.String(http.StatusInternalServerError, "Server Error")
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status": "ok",
	})
}

// タスクの削除
func TaskDelete(c *gin.Context) {

	// taskidをURLから抽出
	taskid, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.String(http.StatusBadRequest, "Bad request")
		return
	}

	task := model.Task{
		TaskID: taskid,
	}

	TaskService := service.TaskService{}

	// 削除関数を呼び出し
	if err := TaskService.DeleteTask(&task); err != nil {
		c.String(http.StatusInternalServerError, "Server Error")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}
