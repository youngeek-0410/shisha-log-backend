package handler

import (
	"net/http"
	"shisha-log-backend/model/flavor"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func UserFlavorResponseGet(GetUserFlavorResponse *flavor.GetUserFlavorResponse) gin.HandlerFunc {
	return func(c *gin.Context) {
		result, err := GetUserFlavorResponse.GetUserFlavorResponse(c.Param("user_id"))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, result)
	}
}

func DiaryFlavorsGet(diaryFlavors *flavor.DiaryFlavors) gin.HandlerFunc {
	return func(c *gin.Context) {
		result, err := diaryFlavors.DiaryFlavors(c.Param("diary_id"))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, result)
	}
}

func CreateFlavorBrand(flavorBrand *flavor.FlavorBrand) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req flavor.FlavorBrand
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		brandID, err := uuid.New().MarshalBinary()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		flavorBrandItem := flavor.FlavorBrand{
			ID:        brandID,
			Name:      req.Name,
			CreatedAt: req.CreatedAt,
		}

		err = flavorBrand.Add(flavorBrandItem)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Successful operation"})
	}
}
