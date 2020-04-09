package routers

import (
	"github.com/gin-gonic/gin"
	"go_cron/controllers"
	jwt "go_cron/middleware"
	"go_cron/pkg/setting"
	"go_cron/routers/v1"
)

func InitRouter() *gin.Engine {

	r := gin.Default()

	r.POST("/login", controllers.GetAuth) // 认证
	gin.SetMode(setting.RunMode)

	//路由组
	user := r.Group("/user") // users/
	user.Use(jwt.JWT())      // 加载认证中间件

	v1.UserRouter(user) // 用户路由

	//路由组
	task := r.Group("/task")
	v1.TaskRouter(task)

	//路由组
	taskLog := r.Group("/taskLog")
	v1.TaskLogRouter(taskLog)

	//路由组
	taskGroup := r.Group("/taskGroup")
	v1.TaskGroupRouter(taskGroup)

	r.POST("/home/dashboard", jwt.JWT(), controllers.DashBoard) // 需要认证

	return r

}
