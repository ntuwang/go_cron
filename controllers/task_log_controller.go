package controllers

import (
	"github.com/gin-gonic/gin"
	"go_cron/models"
	"go_cron/pkg/e"
	"net/http"
)

func TaskLogList(c *gin.Context) {
	code := e.SUCCESS
	query := models.Query{"", 3, ""}

	err := c.BindJSON(&query)
	if err != nil {
		code = e.INVALID_PARAMS
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  e.GetMsg(code),
			"data": "",
		})
		return
	}

	taskLogs, _ := models.ListTaskLog(query)
	values := make(map[string]interface{})
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": taskLogs,
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
