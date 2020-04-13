package models

import (
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"strings"
)

type TaskLog struct {
	Id          int    `gorm:"size:10;primary_key;AUTO_INCREMENT;not null" json:"id"`
	TaskId      int    `gorm:"size:10;DEFAULT NULL" json:"taskId"`
	Task        Task   `gorm:"foreignkey:TaskId" json:"task"`
	Output      string `gorm:"size:200;DEFAULT NULL" json:"output"`
	Error       string `gorm:"size:200;DEFAULT NULL" json:"error"`
	Status      int    `gorm:"size:10;DEFAULT NULL" json:"status"`
	ProcessTime int    `gorm:"size:64;DEFAULT NULL" json:"processTime"`
	CreateTime  string `gorm:"size:50;DEFAULT NULL" json:"createTime"`
}

type Query struct {
	TaskName string `json:"taskName"`
	Status   int    `json:"status"`
	Datetime string `json:"datetime"`
}

func (TaskLog) TableName() string {
	return "t_task_log"
}

func CreateTaskLog(taskLog *TaskLog) (int, error) {

	row := new(TaskLog)
	d := db.Create(taskLog).Scan(&row)
	if d.Error != nil {
		return 0, d.Error
	}
	return row.Id, nil

}

func ListTaskLog(query Query) ([]TaskLog, error) {

	var Db = db
	if query.TaskName != "" {
		Db = Db.Where("task_id = ?", query.TaskName)
	}
	if query.Status != 3 {
		Db = Db.Where("status = ?", query.Status)
	}
	if query.Datetime != "" {
		datetime := strings.Split(query.Datetime, "~")
		startTime := strings.Trim(datetime[0], " ")
		endTime := strings.Trim(datetime[1], " ")
		Db = Db.Where("create_time > ?", startTime).Where("create_time < ?", endTime)
	}
	var taskLog []TaskLog
	err := Db.Preload("Task").Order("id desc").Find(&taskLog).Error

	return taskLog, err
}

func GetTaskLogByTaskId(taskId int) ([]TaskLog, error) {

	//defer db.Close()
	var taskLog []TaskLog
	err := db.Where("task_id = ?", taskId).Find(&taskLog).Error
	fmt.Println(err, taskLog)
	return taskLog, err
}

func DeleteTaskLogByTaskId(logName string) bool {

	var taskLog TaskLog
	err := db.Where("task_id = ?", logName).Delete(&taskLog).Error
	if err != nil {
		return false
	}
	return true

}

func ListSuccessTask() ([]TaskLog, error) {

	var t []TaskLog
	//err := db.Limit(3).Find(&task).Error //限制查找前line行
	err := db.Where("status = 0").Find(&t).Error

	return t, err
}

func ListFailedTask() ([]TaskLog, error) {

	var t []TaskLog
	//err := db.Limit(3).Find(&task).Error //限制查找前line行
	err := db.Where("status <> 0").Find(&t).Error

	return t, err
}
