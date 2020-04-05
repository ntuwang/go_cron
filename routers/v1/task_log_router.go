package v1

import (
	"github.com/gin-gonic/gin"
	"go_cron/controllers"
)

func TaskLogRouter(task *gin.RouterGroup) {
	task.GET(`/TaskLogList`, controllers.TaskLogList)      // 判断账号是否存在
	task.POST(`/TaskLogAdd`, controllers.TaskLogAdd)       // 登陆
	task.GET(`/TaskLogInfo`, controllers.TaskLogInfo)      // 登陆
	task.POST(`/TaskLogDelete`, controllers.TaskLogDelete) // 登陆
}
