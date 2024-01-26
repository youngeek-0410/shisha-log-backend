package handler

import (
	"net/http"
	"shisha-log-backend/model/bottle"

	"github.com/gin-gonic/gin"
)

// func BottlesGet(bottles *bottle.Bottles) gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		result := bottles.GetAll()
// 		c.JSON(http.StatusOK, result)
// 	}
// }

func UserBottlesGet(userBottles *bottle.UserBottles) gin.HandlerFunc {
	return func(c *gin.Context) {
		result, err := userBottles.UserBottles(c.Param("user_id"))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, result)
	}
}