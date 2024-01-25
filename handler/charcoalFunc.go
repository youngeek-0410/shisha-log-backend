package handler

import (
	"net/http"
	"shisha-log-backend/model/charcoal"

	"github.com/gin-gonic/gin"
)

func UserCharcoalsGet(userCharcoals *charcoal.UserCharcoals) gin.HandlerFunc {
	return func(c *gin.Context) {
		result, err := userCharcoals.UserCharcoals(c.Param("user_id"))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, result)
	}
}
