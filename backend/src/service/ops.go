package service

import (
	"iotproject/db"
	"iotproject/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func getByID(ctx *gin.Context) {
	cardID := ctx.Query("cardID")
	if cardID == "" {
		cardID = ctx.PostForm("cardID")
	}
	ret := &model.Permission{}
	if err := db.GetDB().First(ret, "card_id = ?", cardID).Error; err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"username":   "unknown",
			"cardID":     cardID,
			"permission": "no",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"username":   ret.Username,
		"cardID":     cardID,
		"permission": "yes",
	})
}

func getByUser(ctx *gin.Context) {
	username := ctx.Query("username")
	if username == "" {
		username = ctx.PostForm("username")
	}
	ret := &model.Permission{}
	if err := db.GetDB().First(ret, "username = ?", username).Error; err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"username":   username,
			"cardID":     "unknown",
			"permission": "no",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"username":   ret.Username,
		"cardID":     ret.CardID,
		"permission": "yes",
	})
}

// Get permission by card id
func GetByID(ctx *gin.Context) {
	getByID(ctx)
}

func GetByUser(ctx *gin.Context) {
	getByUser(ctx)
}

// add username and card id to database
func AddUserID(ctx *gin.Context) {
	cardID := ctx.PostForm("cardID")
	username := ctx.PostForm("username")

	if err := db.GetDB().Create(
		&model.Permission{
			CardID:   cardID,
			Username: username,
		}).Error; err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"username":   username,
			"cardID":     cardID,
			"permission": "no",
		})
		return
	}

	getByID(ctx)
}

// delete by card id
func DeleteByID(ctx *gin.Context) {
	cardID := ctx.PostForm("cardID")

	db.GetDB().Delete(&model.Permission{}, "card_id = ?", cardID)
	getByID(ctx)
}

func DeleteByUser(ctx *gin.Context) {
	username := ctx.PostForm("username")

	db.GetDB().Delete(&model.Permission{}, "username = ?", username)
	getByUser(ctx)
}
