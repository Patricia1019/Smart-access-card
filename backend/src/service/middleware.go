package service

import (
	"iotproject/db"
	"iotproject/model"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func Verify(ctx *gin.Context) {
	// verify... except for /api/getByID
	// TODO(yuanyanlong): figure out a more elegant way to do this
	if ctx.Request.URL.Path == "/api/getByID" {
		ctx.Next()
		return
	}
	if cookie, err := ctx.Cookie("auth"); err == nil {
		auth := &model.UserToken{}
		if err := db.GetDB().First(auth, "token = ?", cookie).Error; err == nil {
			// fmt.Println(cookie, auth.Token)
			if auth.ExpiredAt > time.Now().Unix() {
				ctx.Next()
				return
			}
		}
	}
	ctx.HTML(http.StatusUnauthorized, "unauth.html", gin.H{})
	ctx.Abort()
}
