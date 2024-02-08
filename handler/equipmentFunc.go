package handler

import (
	"net/http"
	"shisha-log-backend/model/equipment/bottle"
	"shisha-log-backend/model/equipment/bowl"
	"shisha-log-backend/model/equipment/charcoal"
	"shisha-log-backend/model/equipment/heatmanagement"
	"shisha-log-backend/model/flavor"

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

		flavors, flavorErr := userFlavors.UserFlavors(c.Param("user_id"))
		if flavorErr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": flavorErr.Error()})
			return
		}
		result.Flavors = flavors

		bottles, bottleErr := userBottles.UserBottles(c.Param("user_id"))
		if bottleErr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": bottleErr.Error()})
			return
		}
		result.Bottles = bottles

		bowls, bowlErr := userBowls.UserBowls(c.Param("user_id"))
		if bowlErr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": bowlErr.Error()})
			return
		}
		result.Bowls = bowls

		charcoals, charcoalErr := userCharcoals.UserCharcoals(c.Param("user_id"))
		if charcoalErr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": charcoalErr.Error()})
			return
		}
		result.Charcoals = charcoals

		heatManagements, heatManagementErr := userHeatManagements.UserHeatManagements(c.Param("user_id"))
		if heatManagementErr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": heatManagementErr.Error()})
			return
		}
		result.HeatManagements = heatManagements

		c.JSON(http.StatusOK, result)
	}
}
