package main

import (
	"iotproject/db"
	"iotproject/model"
	"iotproject/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

var HttpServer http.Server

func main() {
	db.InitDB("mysql", "root:ABhO3WZV@tcp(localhost:3306)/mysql?charset=utf8mb4&parseTime=True&loc=Local")
	db.AutoMigrate(&model.Permission{}, &model.Admin{}, &model.UserToken{})

	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.Static("../static", "static")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	// login page
	r.POST("/login", service.Login)

	// operations
	op := r.Group("api")
	// register middleware
	op.Use(service.Verify)
	{
		op.GET("/getByID", service.GetByID)
		op.GET("/getByUser", service.GetByUser)
		op.POST("/addUserID", service.AddUserID)
		op.POST("/deleteByID", service.DeleteByID)
		op.POST("/deleteByUser", service.DeleteByUser)
	}
	r.Run(":8080")
}
