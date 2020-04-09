package controllers

import (
	"github.com/gin-gonic/gin"
	"go_cron/models"
	"go_cron/pkg/e"
	"net/http"
)

func DashBoard(c *gin.Context) {
	//返回JSON
	code := 200
	users, _ := models.ListUser()
	tasks, _ := models.ListTask()

	var tLogs []interface{}
	for _, t := range tasks {
		d := make(map[string]interface{})
		d["taskName"] = t.TaskName
		taskLog, _ := models.GetTaskLogByTaskId(t.Id)
		m := taskLog[len(taskLog)-1]
		d["taskLog"] = m
		tLogs = append(tLogs, d)
	}
	successTask, _ := models.ListSuccessTask()
	failedTask, _ := models.ListFailedTask()
	data := make(map[string]interface{})
	data["userLength"] = len(users)
	data["taskLength"] = len(tasks)
	data["successTaskLength"] = len(successTask)
	data["failedTaskLength"] = len(failedTask)
	data["tLogs"] = tLogs

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}
