package v1

import (
	"github.com/gin-gonic/gin"
	"go_cron/controllers"
)

func TaskGroupRouter(task *gin.RouterGroup) {
	task.GET(`/TaskGroupList`, controllers.TaskGroupList)      // 判断账号是否存在
	task.POST(`/TaskGroupAdd`, controllers.TaskGroupAdd)       // 登陆
	task.GET(`/TaskGroupInfo`, controllers.TaskGroupInfo)      // 登陆
	task.POST(`/TaskGroupDelete`, controllers.TaskGroupDelete) // 登陆
}
