package models

import (
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

const (
	TASK_SUCCESS = 0  // 任务执行成功
	TASK_ERROR   = -1 // 任务执行出错
	TASK_TIMEOUT = -2 // 任务执行超时
)

type Task struct {
	Id           int       `gorm:"size:10;primary_key;AUTO_INCREMENT;not null" json:"id"`
	UserId       int       `gorm:"size:20;DEFAULT NULL" json:"userId"`
	User         User      `json:"user"`
	GroupId      int       `gorm:"index" json:"groupId"`
	TaskGroup    TaskGroup `gorm:"foreignkey:GroupId" json:"taskGroup"`
	TaskName     string    `gorm:"size:20;DEFAULT NULL" json:"taskName"`
	TaskType     int       `gorm:"size:20;DEFAULT NULL" json:"taskType"`
	Description  string    `gorm:"size:50;DEFAULT NULL" json:"description"`
	CronSpec     string    `gorm:"size:20;DEFAULT NULL" json:"cronSpec"`
	Concurrent   int       `gorm:"size:20;DEFAULT NULL" json:"concurrent"`
	Command      string    `gorm:"size:200;DEFAULT NULL" json:"command"`
	Status       int       `gorm:"size:20;DEFAULT NULL" json:"status"`
	Notify       int       `gorm:"size:20;DEFAULT NULL" json:"notify"`
	NotifyEmail  string    `gorm:"size:200;DEFAULT NULL" json:"notifyEmail"`
	Timeout      int       `gorm:"size:20;DEFAULT NULL" json:"timeout"`
	ExecuteTimes int       `gorm:"size:20;DEFAULT NULL" json:"executeTimes"`
	PrevTime     string    `gorm:"size:50;DEFAULT NULL" json:"prevTime"`
	CreateTime   string    `gorm:"size:50;DEFAULT NULL" json:"createTime"`
}

func (Task) TableName() string {
	return "t_task"
}

func CreateTask(task *Task) bool {
	t := time.Now().Format("2006-01-02 15:04:05") //2019-07-31 13:55:21.3410012 +0800 CST m=+0.006015601
	task.CreateTime = t
	err := db.Create(&task).Error //创建对象
	if err != nil {
		return false
	}
	return true
}

func ListTask() ([]Task, error) {

	var tasks []Task
	//err := db.Limit(3).Find(&task).Error //限制查找前line行
	err := db.Preload("TaskGroup").Find(&tasks).Error

	return tasks, err
}

func ListTaskTotal() ([]Task, int) {

	var tasks []Task
	//err := db.Limit(3).Find(&task).Error //限制查找前line行
	db.Preload("TaskGroup").Find(&tasks)
	total := len(tasks)

	return tasks, total
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

func UpdateTask(taskName string, values map[string]interface{}) bool {
	var task Task
	db.Select("id").Where("task_name = ?", taskName).First(&task)
	if task.Id == 0 {
		return false
	}
	err := db.Model(&task).Updates(values).Error
	if err != nil {
		return false
	}
	return true
}
