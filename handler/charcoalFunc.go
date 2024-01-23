package handler

import (
	"net/http"
	"shisha-log-backend/model/charcoal"

	"github.com/gin-gonic/gin"
)

func UserCharcoalsGet(userCharcoals *charcoal.UserCharcoals) gin.HandlerFunc {
	return func(c *gin.Context) {
		result := userCharcoals.UserCharcoals(c.Param("user_id"))
		c.JSON(http.StatusOK, result)
	}
}
