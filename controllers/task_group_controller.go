package controllers

import (
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

	//返回JSON
	user := &models.TaskGroup{}
	err := c.Bind(user)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 0, "message": err})
		return
	} else {
		res := models.CreateTaskGroup(user)
		c.JSON(http.StatusOK, gin.H{"code": 0, "message": res})
	}
}

func TaskGroupInfo(c *gin.Context) {
	groupName := c.Query("groupName")

	user, _ := models.GetTaskGroupByGroupName(groupName)
	c.JSON(http.StatusOK, user)
}

func TaskGroupDelete(c *gin.Context) {
	groupName := c.PostForm("groupName")

	res := models.DeleteTaskGroupByGroupName(groupName)
	c.JSON(http.StatusOK, res)
}
