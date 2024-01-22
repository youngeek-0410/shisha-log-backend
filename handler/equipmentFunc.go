package handler

import (
	"net/http"
	"shisha-log-backend/model/bottle"
	"shisha-log-backend/model/bowl"

	"github.com/gin-gonic/gin"
)

type UserEquipments struct {
	Bottles []bottle.UserBottle `json:"user_bottle_list"`
	Bowls   []bowl.UserBowl     `json:"user_bowl_list"`
}

func NewUserEquipments() *UserEquipments {
	return &UserEquipments{}
}

func UserEquipmentsGet(UserEquipment *UserEquipments) gin.HandlerFunc {
	userBottles := bottle.NewUserBottles()
	userBowls := bowl.NewUserBowls()

	return func(c *gin.Context) {
		var result UserEquipments
		result.Bottles = userBottles.UserBottles(c.Param("user_id"))
		result.Bowls = userBowls.UserBowls(c.Param("user_id"))
		c.JSON(http.StatusOK, result)
	}
}
