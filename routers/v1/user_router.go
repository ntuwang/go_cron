package v1

import (
	"github.com/gin-gonic/gin"
	"go_cron/controllers"
)

func UserRouter(user *gin.RouterGroup) {
	user.GET(`/UserList`, controllers.UserList)              // 判断账号是否存在
	user.POST(`/UserAdd`, controllers.UserAdd)               // 登陆
	user.POST(`/ChangePassword`, controllers.ChangePassword) // 登陆
	user.GET(`/UserInfo`, controllers.UserInfo)              // 登陆
	user.GET(`/LoginUser`, controllers.UserInfoByToken)      // 登陆
	user.POST(`/UserDelete`, controllers.UserDelete)         // 登陆
}
