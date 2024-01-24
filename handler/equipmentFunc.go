package handler

import (
	"net/http"
	"shisha-log-backend/model/bottle"
	"shisha-log-backend/model/bowl"
	"shisha-log-backend/model/charcoal"
	"shisha-log-backend/model/flavor"
	"shisha-log-backend/model/heatmanagement"

	"github.com/gin-gonic/gin"
)

type UserEquipments struct {
	Flavors         []flavor.UserFlavor                 `json:"user_flavor_list"`
	Bottles         []bottle.UserBottle                 `json:"user_bottle_list"`
	Bowls           []bowl.UserBowl                     `json:"user_bowl_list"`
	Charcoals       []charcoal.UserCharcoal             `json:"user_charcoal_list"`
	HeatManagements []heatmanagement.UserHeatManagement `json:"user_heat_management_list"`
}

func NewUserEquipments() *UserEquipments {
	return &UserEquipments{}
}

func UserEquipmentsGet(UserEquipment *UserEquipments) gin.HandlerFunc {
	userFlavors := flavor.NewUserFlavors()
	userBottles := bottle.NewUserBottles()
	userBowls := bowl.NewUserBowls()
	userCharcoals := charcoal.NewUserCharcoals()
	userHeatManagements := heatmanagement.NewUserHeatManagements()

	return func(c *gin.Context) {
		var result UserEquipments
		result.Flavors = userFlavors.UserFlavors(c.Param("user_id"))
		result.Bottles = userBottles.UserBottles(c.Param("user_id"))
		result.Bowls = userBowls.UserBowls(c.Param("user_id"))
		result.Charcoals = userCharcoals.UserCharcoals(c.Param("user_id"))
		result.HeatManagements = userHeatManagements.UserHeatManagements(c.Param("user_id"))
		c.JSON(http.StatusOK, result)
	}
}
