package models

import (
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Task struct {
	Id           int    `gorm:"size:10;primary_key;AUTO_INCREMENT;not null" json:"id"`
	UserId       int    `gorm:"size:20;DEFAULT NULL" json:"userId"`
	GroupId      int    `gorm:"size:20;DEFAULT NULL" json:"groupId"`
	TaskName     string `gorm:"size:20;DEFAULT NULL" json:"taskName"`
	TaskType     int    `gorm:"size:20;DEFAULT NULL" json:"taskType"`
	Description  string `gorm:"size:50;DEFAULT NULL" json:"description"`
	CronSpec     string `gorm:"size:20;DEFAULT NULL" json:"cronSpec"`
	Concurrent   int    `gorm:"size:20;DEFAULT NULL" json:"concurrent"`
	Command      string `gorm:"size:200;DEFAULT NULL" json:"command"`
	Status       int    `gorm:"size:20;DEFAULT NULL" json:"status"`
	Notify       int    `gorm:"size:20;DEFAULT NULL" json:"notify"`
	NotifyEmail  string `gorm:"size:200;DEFAULT NULL" json:"notifyEmail"`
	Timeout      int    `gorm:"size:20;DEFAULT NULL" json:"timeout"`
	ExecuteTimes int    `gorm:"size:20;DEFAULT NULL" json:"executeTimes"`
	PrevTime     int64  `gorm:"size:64;DEFAULT NULL" json:"prevTime"`
	CreateTime   int    `gorm:"size:64;DEFAULT NULL" json:"createTime"`
}

func (Task) TableName() string {
	return "t_task"
}

func CreateTask(task *Task) bool {

	err := db.Create(&task).Error //创建对象
	if err != nil {
		return false
	}
	return true
}

func ListTask() ([]Task, error) {

	var task []Task
	//err := db.Limit(3).Find(&task).Error //限制查找前line行
	err := db.Find(&task).Error

	return task, err
}

func GetTaskByTaskName(taskName string) (Task, error) {

	//defer db.Close()
	var task Task
	err := db.Debug().Where("task_name = ?", taskName).First(&task).Error
	fmt.Println(err, task)
	return task, err
}

func DeleteTaskByTaskName(taskName string) bool {

	var task Task
	err := db.Where("task_name = ?", taskName).Delete(&task).Error
	if err != nil {
		return false
	}
	return true

}

func ExistTaskByTaskName(taskName string) bool {
	var task Task
	db.Select("id").Where("task_name = ?", taskName).First(&task)
	if task.Id > 0 {
		return true
	}

	return false
}
