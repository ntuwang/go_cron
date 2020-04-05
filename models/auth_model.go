package models

import "fmt"

func CheckAuth(username, password string) bool {
	var auth User
	fmt.Println(7, username, password)
	db.Select("id").Where(User{Username: username, Password: password}).First(&auth)
	if auth.Id > 0 {
		return true
	}
	return false
}
