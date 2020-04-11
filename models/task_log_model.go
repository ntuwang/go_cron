package models

import (
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type TaskLog struct {
	Id          int    `gorm:"size:10;primary_key;AUTO_INCREMENT;not null" json:"id"`
	TaskId      int    `gorm:"size:10;DEFAULT NULL" json:"taskId"`
	Output      string `gorm:"size:200;DEFAULT NULL" json:"output"`
	Error       string `gorm:"size:200;DEFAULT NULL" json:"error"`
	Status      int    `gorm:"size:10;DEFAULT NULL" json:"status"`
	ProcessTime int    `gorm:"size:64;DEFAULT NULL" json:"processTime"`
	CreateTime  string `gorm:"size:50;DEFAULT NULL" json:"createTime"`
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

func ListTaskLog() ([]TaskLog, error) {

	var taskLog []TaskLog
	//err := db.Limit(3).Find(&taskLog).Error //限制查找前line行
	err := db.Find(&taskLog).Error

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
