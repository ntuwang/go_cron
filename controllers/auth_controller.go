package controllers

import (
	"fmt"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"

	"go_cron/models"
	"go_cron/pkg/e"
	"go_cron/pkg/util"
)

type User struct {
	Username string `gorm:"size:20;DEFAULT NULL" json:"username" valid:"Required; MaxSize(50)"`
	Password string `gorm:"size:100;DEFAULT NULL" json:"password" valid:"Required; MaxSize(100)"`
}

func GetAuth(c *gin.Context) {

	a := User{}
	c.BindJSON(&a)
	fmt.Println(a)
	username := a.Username
	password := util.MD5(c.PostForm(a.Password))

	valid := validation.Validation{}
	//a := User{Username: username, Password: password}
	ok, _ := valid.Valid(&a)

	data := make(map[string]interface{})

	nowTime := time.Now()
	data["expireTime"] = nowTime.Add(8 * time.Hour).Unix() // 得到秒级时间戳
	code := e.INVALID_PARAMS
	if ok {
		isExist := models.CheckAuth(username, password)
		if isExist {
			token, err := util.GenerateToken(username, password)
			if err != nil {
				code = e.ERROR_AUTH_TOKEN
			} else {
				data["token"] = token

				code = e.SUCCESS
			}

		} else {
			code = e.ERROR_LOGIN
		}
	} else {
		fmt.Println(username, password, 11111111)
		for _, err := range valid.Errors {
			log.Println(err.Key, err.Message)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

// 根据token值解析用户信息
func UserInfoByToken(c *gin.Context) {
	var code int
	code = e.SUCCESS
	token := c.Request.Header.Get("Authorization")
	claims, err := util.ParseToken(token)
	fmt.Println(claims, "登录的用户信息")
	if err != nil {
		code = e.ERROR_AUTH_CHECK_TOKEN_FAIL
	} else if time.Now().Unix() > claims.ExpiresAt {
		code = e.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
	} else {
		if code == e.SUCCESS {
			user, _ := models.GetUserByUsername(claims.Username)
			c.JSON(http.StatusOK, gin.H{
				"code": code,
				"msg":  e.GetMsg(code),
				"data": user,
			})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"code": code, "msg": e.GetMsg(code)})
}
