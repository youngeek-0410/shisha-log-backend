package handler

import (
	"net/http"
	"shisha-log-backend/model/equipment/heatmanagement"

	"github.com/gin-gonic/gin"
)

func UserHeatManagementsGet(userHeatManagements *heatmanagement.UserHeatManagements) gin.HandlerFunc {
	return func(c *gin.Context) {
		result, err := userHeatManagements.UserHeatManagements(c.Param("user_id"))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, result)
	}
}
