package handler

import (
	"net/http"
	"shisha-log-backend/model/flavor"

	"github.com/gin-gonic/gin"
)

func UserFlavorsGet(userFlavors *flavor.UserFlavors) gin.HandlerFunc {
	return func(c *gin.Context) {
		result, err := userFlavors.UserFlavors(c.Param("user_id"))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, result)
	}
}

func DiaryFlavorsGet(diaryFlavors *flavor.DiaryFlavors) gin.HandlerFunc {
	return func(c *gin.Context) {
		result, err := diaryFlavors.DiaryFlavors(c.Param("diary_id"))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, result)
	}
}
