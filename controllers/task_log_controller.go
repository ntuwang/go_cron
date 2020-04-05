package controllers

import (
	"github.com/gin-gonic/gin"
	"go_cron/models"
	"net/http"
)

func TaskLogList(c *gin.Context) {
	//返回JSON
	user, _ := models.ListTaskLog()
	c.JSON(http.StatusOK, user)
}

func TaskLogAdd(c *gin.Context) {

	//返回JSON
	user := &models.TaskLog{}
	err := c.Bind(user)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 0, "message": err})
		return
	} else {
		res := models.CreateTaskLog(user)
		c.JSON(http.StatusOK, gin.H{"code": 0, "message": res})
	}
}

func TaskLogInfo(c *gin.Context) {
	taskId := c.Query("taskId")

	user, _ := models.GetTaskLogByTaskId(taskId)
	c.JSON(http.StatusOK, user)
}

func TaskLogDelete(c *gin.Context) {
	taskId := c.PostForm("taskId")

	res := models.DeleteTaskLogByTaskId(taskId)
	c.JSON(http.StatusOK, res)
}
