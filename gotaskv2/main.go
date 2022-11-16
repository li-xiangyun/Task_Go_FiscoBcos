package main

import (
	"net/http"

	"gotask/routes"

	"github.com/gin-gonic/gin"
	"github.com/go-session/gin-session"
)

func main() {

	//创建gin对象
	r := gin.Default()
	r.Use(ginsession.New())
	//设置静态页面
	r.StaticFile("/", "./dist/index.html")
	r.StaticFile("/index", "./dist/index.html")
	r.Static("/static", "./dist/static")

	//设置路由规则
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	//注册规则
	r.POST("/register", routes.Register)
	r.POST("/login", routes.Login)
	//任务发布
	r.POST("/issue", routes.Issue)
	//任务修改
	r.POST("/update", routes.Update)
	//任务查询
	r.GET("/tasklist", routes.Tasklist)
	//启动web服务
	r.Run(":9090") // listen and serve on 0.0.0.0:8080 (for windows "localhost:9090")
}
