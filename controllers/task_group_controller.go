package controllers

import (
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"go_cron/models"
	"go_cron/pkg/e"
	"net/http"
)

func TaskGroupList(c *gin.Context) {
	//返回JSON
	code := 200
	taskGroup, _ := models.ListTaskGroup()
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": taskGroup,
	})
}

func TaskGroupAdd(c *gin.Context) {

	taskGroup := &models.TaskGroup{}
	err := c.Bind(taskGroup)
	code := e.SUCCESS
	if err != nil {
		code = e.INVALID_PARAMS
	} else {

		valid := validation.Validation{}
		valid.Required(taskGroup.GroupName, "groupName").Message("名称不能为空")

		if valid.HasErrors() {
			code = e.INVALID_PARAMS
		} else {
			isExist := models.ExistTaskGroupByGroupName(taskGroup.GroupName)
			if isExist {
				code = e.ERROR_EXIST_TASKGROUP
			} else {
				models.CreateTaskGroup(taskGroup)
			}
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": "",
	})
}

func TaskGroupInfo(c *gin.Context) {
	groupName := c.Query("groupName")

	user, _ := models.GetTaskGroupByGroupName(groupName)
	c.JSON(http.StatusOK, user)
}

func TaskGroupDelete(c *gin.Context) {

	groupName := &models.TaskGroup{}
	err := c.Bind(groupName)
	code := e.SUCCESS
	if err != nil {
		code = e.INVALID_PARAMS
	} else {

		valid := validation.Validation{}
		valid.Required(groupName.GroupName, "groupName").Message("名称不能为空")

		if valid.HasErrors() {
			code = e.INVALID_PARAMS
		} else {
			isExist := models.ExistTaskGroupByGroupName(groupName.GroupName)
			if !isExist {
				code = e.ERROR_NOT_EXIST_TASKGROUP
			} else {
				models.DeleteTaskGroupByGroupName(groupName.GroupName)
			}
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": "",
	})
}
