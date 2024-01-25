package handler

import (
	"net/http"
	"shisha-log-backend/model/bowl"

	"github.com/gin-gonic/gin"
)

func BowlsGet(bowls *bowl.Bowls) gin.HandlerFunc {
	return func(c *gin.Context) {
		result := bowls.GetAll()
		c.JSON(http.StatusOK, result)
	}
}

func UserBowlsGet(userBowls *bowl.UserBowls) gin.HandlerFunc {
	return func(c *gin.Context) {
		result, err := userBowls.UserBowls(c.Param("user_id"))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, result)
	}
}
