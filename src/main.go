package main

import (
	"todo.app/controller"

	"github.com/gin-gonic/gin"
)

/*

func setTasktoList(task_id int, name string, msg string, done bool, user_id int) (r *Task) {
	r = new(model.Task)
	r.TaskID = task_id
	r.Name = name
	r.Message = msg
	r.Done = done
	r.UserID = user_id
	return r
}
*/

//type TaskList = []*Task

func setupRouter() *gin.Engine {

	//	var tasks TaskList
	//var task = model.Task{}

	//	tasks = append(tasks, setTasktoList(1, "これはテストです1", "ああああああ", false, 1))
	//	tasks = append(tasks, setTasktoList(2, "これはテストです2", "ああああああ", true, 1))
	//	tasks = append(tasks, setTasktoList(3, "これはテストです3", "ああああああ", true, 1))

	r := gin.Default()
	// TODO: auth
	/*
		var user = model.User{
			Name:     c.PostForm("username"),
			Password: c.PostForm("password"),
		}
	*/
	r.Static("/views", "../views")
	//r.LoadHTMLGlob("../views/static/*.html")

	v1 := r.Group("/api/v1")
	{
		task_api := v1.Group("/task")
		{
			task_api.POST("/add", controller.TaskAdd)
			task_api.GET("/list", controller.TaskList)
			task_api.PUT("/update", controller.TaskUpdate)
			task_api.DELETE("/delete", controller.TaskDelete)
		}
	}
	/*
		r.GET("/", func(c *gin.Context) {
			c.HTML(http.StatusOK, "index.html", gin.H{})
		})
	*/
	r.StaticFile("/mypage", "../views/static/mypage.html")
	/*
		r.POST("/mypage", func(c *gin.Context) {
			c.HTML(http.StatusOK, "mypage.html", gin.H{
				//"username": user.Name,
				//"tasks":    db.Find(&task),
			})
		})
	*/
	return r
}

func main() {
	r := setupRouter()
	r.Run("127.0.0.1:8080")
}
