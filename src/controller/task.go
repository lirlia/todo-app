package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"todo.app/model"
	"todo.app/service"
)

type TaskPostRequest struct {
	Title string `form:"title" binding:"required,max=100"`
	// https://github.com/gin-gonic/gin/issues/814
	Done    *bool  `form:"done" binding:"required"`
	Message string `form:"message" binding:"required,max=1000"`
	UserID  int    `form:"userid"`
}

func TaskAdd(c *gin.Context) {

	// Form(Task)の構造体を用意
	taskReq := TaskPostRequest{}

	// ポインタを渡してリクエスト内容をバインドする
	if err := c.Bind(&taskReq); err != nil {
		c.String(http.StatusBadRequest, "Bad request")
		fmt.Println(err)
		return
	}

	// DBへ挿入
	task := model.Task{
		Title:   taskReq.Title,
		Done:    *taskReq.Done,
		Message: taskReq.Message,
		UserID:  0,
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

func TaskUpdate(c *gin.Context) {
	Task := model.Task{}
	err := c.Bind(&Task)
	if err != nil {
		c.String(http.StatusBadRequest, "Bad request")
		return
	}
	TaskService := service.TaskService{}
	err = TaskService.UpdateTask(&Task)
	if err != nil {
		c.String(http.StatusInternalServerError, "Server Error")
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"status": "ok",
	})
}

func TaskDelete(c *gin.Context) {
	id := c.PostForm("id")
	intId, err := strconv.ParseInt(id, 10, 0)
	if err != nil {
		c.String(http.StatusBadRequest, "Bad request")
		return
	}
	TaskService := service.TaskService{}
	err = TaskService.DeleteTask(int(intId))
	if err != nil {
		c.String(http.StatusInternalServerError, "Server Error")
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"status": "ok",
	})
}
