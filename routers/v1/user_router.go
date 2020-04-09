package v1

import (
	"github.com/gin-gonic/gin"
	"go_cron/controllers"
)

func UserRouter(user *gin.RouterGroup) {
	user.GET(`/UserList`, controllers.UserList)
	user.POST(`/UserAdd`, controllers.UserAdd)
	user.POST(`/ChangePassword`, controllers.ChangePassword)
	user.GET(`/UserInfo`, controllers.UserInfo)
	user.GET(`/LoginUser`, controllers.UserInfoByToken)
	user.POST(`/UserDelete`, controllers.UserDelete)
	user.POST(`/UserUpdate`, controllers.UserUpdate)
}
