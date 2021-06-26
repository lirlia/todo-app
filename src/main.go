package main

import (
	"todo.app/controller"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {

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
			task_api.PUT("/update/:id", controller.TaskUpdate)
			task_api.DELETE("/delete/:id", controller.TaskDelete)
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
