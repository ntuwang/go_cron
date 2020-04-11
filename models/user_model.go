package models

import (
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct {
	Id        int    `gorm:"size:10;primary_key;AUTO_INCREMENT;not null" json:"id"`
	Username  string `gorm:"size:20;DEFAULT NULL" json:"username"`
	Password  string `gorm:"size:20;DEFAULT NULL" json:"password"`
	Salt      string `gorm:"size:200;DEFAULT NULL" json:"salt"`
	Email     string `gorm:"size:50;DEFAULT NULL" json:"email"`
	LastLogin string `gorm:"size:50;DEFAULT NULL" json:"lastLogin"`
	LastIp    string `gorm:"size:32;DEFAULT NULL" json:"lastIp"`
	Status    int    `gorm:"size:2;DEFAULT NULL" json:"status"`
}

func (User) TableName() string {
	return "t_user"
}

func CreateUser(user *User) bool {

	err := db.Create(&user).Error //创建对象
	if err != nil {
		return false
	}
	return true
}

func ChangePassword(username string, passsword string) bool {

	user := User{}
	fmt.Println(username, passsword)
	err := db.Where("username = ?", username).First(&user).Error
	if err != nil {
		fmt.Println(err)
		return false
	}
	user.Password = passsword
	db.Save(&user) //提交更改
	return true
}

func ListUser() ([]User, error) {

	var user []User
	//err := db.Limit(3).Find(&user).Error //限制查找前line行
	err := db.Find(&user).Error

	return user, err
}

func GetUserByUsername(username string) (User, error) {

	//defer db.Close()
	var user User
	err := db.Debug().Where("username = ?", username).First(&user).Error
	fmt.Println(err, user)
	return user, err
}

func DeleteUserByUsername(username string) bool {

	var user User
	err := db.Where("username = ?", username).Delete(&user).Error
	if err != nil {
		return false
	}
	return true

}

func ExistUserByID(username string) bool {
	var user User
	db.Select("id").Where("username = ?", username).First(&user)
	if user.Id > 0 {
		return true
	}

	return false
}

func UpdateUser(username string, user *User) bool {
	db.Model(&User{}).Where("username = ?", username).Updates(&user)

	return true
}

/*func UpdateUser(username string, data interface{}) bool {
	data := make(map[string]interface{})
	data["modified_by"] = modifiedBy
	if name != "" {
		data["name"] = name
	}
	if state != -1 {
		data["state"] = state
	}
	db.Model(&User{}).Where("username = ?", username).Updates(data)

	return true
}
*/
