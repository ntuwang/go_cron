package models

import (
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type TaskLog struct {
	Id          int    `gorm:"size:10;primary_key;AUTO_INCREMENT;not null" json:"id"`
	TaskId      int    `gorm:"size:10;DEFAULT NULL" json:"userId"`
	Output      string `gorm:"size:200;DEFAULT NULL" json:"output"`
	Error       string `gorm:"size:200;DEFAULT NULL" json:"error"`
	Status      int    `gorm:"size:10;DEFAULT NULL" json:"status"`
	ProcessTime int    `gorm:"size:64;DEFAULT NULL" json:"processTime"`
	CreateTime  int64  `gorm:"size:64;DEFAULT NULL" json:"createTime"`
}

func (TaskLog) TableName() string {
	return "t_task_log"
}

func CreateTaskLog(taskLog *TaskLog) bool {

	err := db.Create(&taskLog).Error //创建对象
	if err != nil {
		return false
	}
	return true
}

func ListTaskLog() ([]TaskLog, error) {

	var taskLog []TaskLog
	//err := db.Limit(3).Find(&taskLog).Error //限制查找前line行
	err := db.Find(&taskLog).Error

	return taskLog, err
}

func GetTaskLogByTaskId(logName string) (TaskLog, error) {

	//defer db.Close()
	var taskLog TaskLog
	err := db.Debug().Where("task_id = ?", logName).First(&taskLog).Error
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
