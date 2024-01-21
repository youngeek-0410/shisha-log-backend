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
