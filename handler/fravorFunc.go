package handler

import (
	"fmt"
	"net/http"
	"shisha-log-backend/lib"
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
		fmt.Printf("req: %+v\n", req)

		flavorID, err := uuid.New().MarshalBinary()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		userFlavorID, err := uuid.New().MarshalBinary()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		flavorItem := flavor.Flavor{
			ID:         flavorID,
			BrandID:    lib.ParseUUIDStrToBin(req.BrandID),
			Name:       req.Name,
			CreateArea: req.CreateArea,
		}

		userFlavorItem := flavor.UserFlavor{
			ID:       userFlavorID,
			FlavorID: flavorID,
			UserID:   lib.ParseUUIDStrToBin(req.UserID),
		}

		err = flavorInstance.Add(flavorItem)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		err = userFlavorInstance.Add(userFlavorItem)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Successful operation"})
	}
}
