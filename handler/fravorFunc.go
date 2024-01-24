package handler

import (
	"net/http"
	"shisha-log-backend/model/fravor"

	"github.com/gin-gonic/gin"
)

func UserFravorsGet(userFravors *fravor.UserFravors) gin.HandlerFunc {
	return func(c *gin.Context) {
		result := userFravors.UserFravors(c.Param("user_id"))
		c.JSON(http.StatusOK, result)
	}
}
