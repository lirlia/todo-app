package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"todo.app/model"
	"todo.app/service"
)

func TaskAdd(c *gin.Context) {
	Task := model.Task{}
	err := c.Bind(&Task)
	if err != nil {
		c.String(http.StatusBadRequest, "Bad request")
		return
	}
	TaskService := service.TaskService{}
	err = TaskService.SetTask(&Task)
	if err != nil {
		c.String(http.StatusInternalServerError, "Server Error")
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"status": "ok",
	})
}

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
