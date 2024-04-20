package handler

import (
	"fmt"
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

		brandID := uuid.New().String()

		flavorBrandItem := flavor.FlavorBrand{
			ID:        brandID,
			Name:      req.Name,
			CreatedAt: req.CreatedAt,
		}

		if err := flavorBrand.Add(flavorBrandItem); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Successful operation"})
	}
}

func CreateFlavor() gin.HandlerFunc {
	return func(c *gin.Context) {
		var flavorInstance flavor.Flavor
		var userFlavorInstance flavor.UserFlavor
		var req flavor.PostFlavorRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			fmt.Printf("req: %+v\n", req)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		flavorID := uuid.New().String()
		userFlavorID := uuid.New().String()

		flavorItem := flavor.Flavor{
			ID:         flavorID,
			BrandID:    req.BrandID,
			Name:       req.Name,
			CreateArea: req.CreateArea,
		}

		userFlavorItem := flavor.UserFlavor{
			ID:       userFlavorID,
			FlavorID: flavorID,
			UserID:   req.UserID,
		}

		if err := flavorInstance.Add(flavorItem); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if err := userFlavorInstance.Add(userFlavorItem); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Successful operation"})
	}
}
