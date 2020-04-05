package v1

import (
	"github.com/gin-gonic/gin"
	"go_cron/controllers"
)

func TaskRouter(task *gin.RouterGroup) {
	task.GET(`/TaskList`, controllers.TaskList)      // 判断账号是否存在
	task.POST(`/TaskAdd`, controllers.TaskAdd)       // 登陆
	task.GET(`/TaskInfo`, controllers.TaskInfo)      // 登陆
	task.POST(`/TaskDelete`, controllers.TaskDelete) // 登陆
}
