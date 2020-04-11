package models

import (
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type TaskGroup struct {
	Id          int    `gorm:"size:10;primary_key;AUTO_INCREMENT;not null" json:"id"`
	UserId      int    `gorm:"size:20;DEFAULT NULL" json:"userId"`
	GroupName   string `gorm:"size:20;DEFAULT NULL" json:"groupName"`
	Description string `gorm:"size:50;DEFAULT NULL" json:"description"`
	CreateTime  string `gorm:"size:50;DEFAULT NULL" json:"createTime"`
}

func (TaskGroup) TableName() string {
	return "t_task_group"
}

func CreateTaskGroup(taskGroup *TaskGroup) bool {

	err := db.Create(&taskGroup).Error //创建对象
	if err != nil {
		return false
	}
	return true
}

func ListTaskGroup() ([]TaskGroup, error) {

	var taskGroup []TaskGroup
	//err := db.Limit(3).Find(&taskGroup).Error //限制查找前line行
	err := db.Find(&taskGroup).Error

	return taskGroup, err
}

func GetTaskGroupByGroupName(groupName string) (TaskGroup, error) {

	//defer db.Close()
	var taskGroup TaskGroup
	err := db.Debug().Where("group_name = ?", groupName).First(&taskGroup).Error
	fmt.Println(err, taskGroup)
	return taskGroup, err
}

func DeleteTaskGroupByGroupName(groupName string) bool {

	var taskGroup TaskGroup
	err := db.Where("group_name = ?", groupName).Delete(&taskGroup).Error
	if err != nil {
		return false
	}
	return true

}

func ExistTaskGroupByGroupName(groupName string) bool {
	var taskGroup TaskGroup
	db.Select("id").Where("group_name = ?", groupName).First(&taskGroup)
	if taskGroup.Id > 0 {
		return true
	}

	return false
}
