package controllers

import (
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	libcron "github.com/lisijie/cron"
	"go_cron/models"
	"go_cron/pkg/e"
	"net/http"
)

func TaskList(c *gin.Context) {
	//返回JSON
	code := e.SUCCESS
	task, _ := models.ListTask()
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": task,
	})
}

func TaskAdd(c *gin.Context) {
	task := &models.Task{}
	err := c.Bind(task)
	code := e.SUCCESS
	if err != nil {
		code = e.INVALID_PARAMS
	} else {

		valid := validation.Validation{}
		valid.Required(task.TaskName, "taskName").Message("不能为空")
		valid.Required(task.Command, "command").Message("不能为空")
		valid.Required(task.CronSpec, "cronSpec").Message("不能为空")

		if _, err := libcron.Parse(task.CronSpec); err != nil {
			// Cron表达式校验
			code = e.INVALID_PARAMS
			c.JSON(http.StatusOK, gin.H{
				"code": code,
				"msg":  e.GetMsg(code),
				"data": "",
			})
			return
		}

		if valid.HasErrors() {
			code = e.INVALID_PARAMS
		} else {
			isExist := models.ExistTaskByTaskName(task.TaskName)
			if isExist {
				code = e.ERROR_EXIST_TASK
			} else {
				models.CreateTask(task)
			}
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": "",
	})

}

func TaskInfo(c *gin.Context) {
	taskName := c.Query("taskName")

	user, _ := models.GetTaskByTaskName(taskName)
	c.JSON(http.StatusOK, user)
}

func TaskDelete(c *gin.Context) {

	task := &models.Task{}
	err := c.Bind(task)
	code := e.SUCCESS
	if err != nil {
		code = e.INVALID_PARAMS
	} else {

		valid := validation.Validation{}
		valid.Required(task.TaskName, "taskName").Message("名称不能为空")

		if valid.HasErrors() {
			code = e.INVALID_PARAMS
		} else {
			isExist := models.ExistTaskByTaskName(task.TaskName)
			if !isExist {
				code = e.ERROR_NOT_EXIST_TASK
			} else {
				models.DeleteTaskByTaskName(task.TaskName)
			}
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": "",
	})
}
