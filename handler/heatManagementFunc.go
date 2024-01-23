package handler

import (
	"net/http"
	"shisha-log-backend/model/heatmanagement"

	"github.com/gin-gonic/gin"
)

func UserHeatManagementsGet(userHeatManagements *heatmanagement.UserHeatManagements) gin.HandlerFunc {
	return func(c *gin.Context) {
		result := userHeatManagements.UserHeatManagements(c.Param("user_id"))
		c.JSON(http.StatusOK, result)
	}
}
