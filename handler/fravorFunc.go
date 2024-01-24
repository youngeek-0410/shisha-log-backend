package handler

import (
	"net/http"
	"shisha-log-backend/model/flavor"

	"github.com/gin-gonic/gin"
)

func UserFlavorsGet(userFlavors *flavor.UserFlavors) gin.HandlerFunc {
	return func(c *gin.Context) {
		result := userFlavors.UserFlavors(c.Param("user_id"))
		c.JSON(http.StatusOK, result)
	}
}
