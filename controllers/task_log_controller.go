package controllers

import (
	"github.com/gin-gonic/gin"
	"go_cron/models"
	"go_cron/pkg/e"
	"net/http"
)

func TaskLogList(c *gin.Context) {
	//返回JSON
	code := 200
	taskLog, _ := models.ListTaskLog()
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": taskLog,
	})
}

func TaskLogAdd(c *gin.Context) {

	taskLog := &models.TaskLog{}
	err := c.Bind(taskLog)
	code := e.SUCCESS
	if err != nil {
		code = e.INVALID_PARAMS
	} else {
		models.CreateTaskLog(taskLog)
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": "",
	})
}

func TaskLogDelete(c *gin.Context) {
	taskId := c.PostForm("taskId")

	res := models.DeleteTaskLogByTaskId(taskId)
	c.JSON(http.StatusOK, res)
}
