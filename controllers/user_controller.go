package controllers

import (
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"go_cron/models"
	"go_cron/pkg/e"
	"go_cron/pkg/util"
	"net/http"
)

func UserList(c *gin.Context) {
	//返回JSON
	code := 200
	user, _ := models.ListUser()
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": user,
	})
}

func UserAdd(c *gin.Context) {
	user := &models.User{}
	err := c.Bind(user)
	code := e.SUCCESS
	if err != nil {
		code = e.INVALID_PARAMS
	} else {

		valid := validation.Validation{}
		valid.Required(user.Username, "username").Message("名称不能为空")

		if valid.HasErrors() {
			code = e.INVALID_PARAMS
		} else {
			isExist := models.ExistUserByID(user.Username)
			if isExist {
				code = e.ERROR_EXIST_USER
			} else {
				user.Password = util.MD5(user.Password)
				models.CreateUser(user)
			}
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": "",
	})

}

func ChangePassword(c *gin.Context) {

	//返回JSON
	username := c.PostForm("username") //post方法取相应字段
	password := c.PostForm("password") //post方法取相应字段

	res := models.ChangePassword(username, password)
	c.JSON(http.StatusOK, gin.H{"code": 0, "message": res})
}

func UserInfo(c *gin.Context) {
	username := c.Query("username")

	user, _ := models.GetUserByUsername(username)
	c.JSON(http.StatusOK, user)
}

func UserDelete(c *gin.Context) {

	user := &models.User{}
	err := c.Bind(user)
	code := e.SUCCESS
	if err != nil {
		code = e.INVALID_PARAMS
	} else {

		valid := validation.Validation{}
		valid.Required(user.Username, "username").Message("名称不能为空")

		if valid.HasErrors() {
			code = e.INVALID_PARAMS
		} else {
			isExist := models.ExistUserByID(user.Username)
			if !isExist {
				code = e.ERROR_NOT_EXIST_USER
			} else {
				models.DeleteUserByUsername(user.Username)
			}
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": "",
	})
}

func UserUpdate(c *gin.Context) {

	user := &models.User{}
	err := c.Bind(user)
	code := e.SUCCESS
	if err != nil {
		code = e.INVALID_PARAMS
	} else {

		valid := validation.Validation{}
		valid.Required(user.Username, "username").Message("名称不能为空")

		if valid.HasErrors() {
			code = e.INVALID_PARAMS
		} else {
			isExist := models.ExistUserByID(user.Username)
			if !isExist {
				code = e.ERROR_NOT_EXIST_USER
			} else {
				if user.Password != "" {
					user.Password = util.MD5(user.Password)
				}
				models.UpdateUser(user.Username, user)
			}
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": "",
	})

}
