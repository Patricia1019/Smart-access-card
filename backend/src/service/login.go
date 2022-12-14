package service

import (
	"fmt"
	"iotproject/db"
	"iotproject/model"
	"net/http"
	"strconv"
	"time"

	"math/rand"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

)

func Login(ctx *gin.Context) {
	ret := &model.Admin{}
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")
	if err := db.GetDB().First(ret, "username = ? and password = ?", username, password).Error; err != nil {
		ctx.HTML(http.StatusOK, "admin.html", gin.H{
			"Name":   ctx.PostForm("username"),
			"Status": "fail",
		})
		return
	}

	// TODO(yuanyanlong): check secure and httponly in https
	rand.Seed(time.Now().Unix())
	token := strconv.Itoa(int(rand.Int63()))
	ctx.SetCookie("auth", token, 3600, "/", "8.134.190.122", false, true)
	// Add the cookie
	userToken := &model.UserToken{}
	if err := db.GetDB().Delete(userToken, "admin_id = ?", ret.Id).Error; err != nil {
		fmt.Println(err)
		if err != gorm.ErrRecordNotFound {
			ctx.HTML(http.StatusForbidden, "admin.html", gin.H{
				"Name":   ctx.PostForm("username"),
				"Status": "fail",
			})
			return
		}
	}
	db.GetDB().Create(
		&model.UserToken{
			AdminId:    ret.Id,
			Token:      token,
			CreateTime: time.Now(),
			ExpiredAt:  time.Now().Unix() + 3600,
		},
	)
	ctx.HTML(http.StatusOK, "data.html", gin.H{
		"Name":   ctx.PostForm("username"),
		"Status": "pass",
	})
}
