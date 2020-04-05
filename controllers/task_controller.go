package controllers

import (
	"github.com/gin-gonic/gin"
	"go_cron/models"
	"go_cron/pkg/e"
	"net/http"
)

func TaskList(c *gin.Context) {
	//返回JSON
	code := 200
	task, _ := models.ListTask()
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": task,
	})
}

func TaskAdd(c *gin.Context) {

	//返回JSON
	user := &models.Task{}
	err := c.Bind(user)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 0, "message": err})
		return
	} else {
		res := models.CreateTask(user)
		c.JSON(http.StatusOK, gin.H{"code": 0, "message": res})
	}
}

func TaskInfo(c *gin.Context) {
	taskName := c.Query("taskName")

	user, _ := models.GetTaskByTaskName(taskName)
	c.JSON(http.StatusOK, user)
}

func TaskDelete(c *gin.Context) {
	taskName := c.PostForm("taskName")

	res := models.DeleteTaskByTaskName(taskName)
	c.JSON(http.StatusOK, res)
}
