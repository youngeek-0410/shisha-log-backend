package handler

import (
	"net/http"
	"shisha-log-backend/model/bottle"

	"github.com/gin-gonic/gin"
)

func BottlesGet(bottles *bottle.Bottles) gin.HandlerFunc {
	return func(c *gin.Context) {
		result := bottles.GetAll()
		c.JSON(http.StatusOK, result)
	}
}

func UserBottlesGet(userBottles *bottle.UserBottles) gin.HandlerFunc {
	return func(c *gin.Context) {
		result := userBottles.UserBowls(c.Param("user_id"))
		c.JSON(http.StatusOK, result)
	}
}
