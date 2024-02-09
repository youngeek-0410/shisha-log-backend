package handler

import (
	"net/http"
	"shisha-log-backend/lib"
	"shisha-log-backend/model/diary"
	"shisha-log-backend/model/equipment"
	"shisha-log-backend/model/flavor"
	"shisha-log-backend/model/user"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetUserDiaries(userDiaries *diary.UserDiaries) gin.HandlerFunc {
	return func(c *gin.Context) {
		result, err := userDiaries.UserDiaries(c.Param("user_id"))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, result)
	}
}

func CreateDiary(c *gin.Context) {
	var req diary.DiaryRequest
	var diaries diary.Diaries
	var diaryEquipments equipment.DiaryEquipments
	var diaryFlavors flavor.PostDiaryFlavors
	var userDiaries user.UserDiaries

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	diaryID, err := uuid.New().MarshalBinary()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	diaryEquipmentsID, err := uuid.New().MarshalBinary()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	userDiariesID, err := uuid.New().MarshalBinary()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// テストのために一時的にimage_idの外部キー制約外してるから戻す
	diaryEquipmentsItem := equipment.DiaryEquipment{
		ID:                   diaryEquipmentsID,
		UserBowlID:           lib.ParseUUIDStrToBin(req.Equipments.BowlID),
		UserBottleID:         lib.ParseUUIDStrToBin(req.Equipments.BottleID),
		UserHeatManagementID: lib.ParseUUIDStrToBin(req.Equipments.HeatManagementID),
		UserCharcoalID:       lib.ParseUUIDStrToBin(req.Equipments.CharcoalID),
		DiaryImageID:         lib.ParseUUIDStrToBin(req.ImageID),
	}

	diaryItem := diary.Diary{
		ID:                diaryID,
		DiaryEquipmentsID: diaryEquipmentsID,
		SuckingText:       &req.SuckingText,
		Temperature:       &req.Equipments.Climate.Temperature,
		Humidity:          &req.Equipments.Climate.Humidity,
		CreatorEvaluation: req.Review.CreatorEvaluation,
		TasteEvaluation:   req.Review.TasteEvaluation,
		CreatorGoodPoints: &req.Review.CreatorGoodPoints,
		CreatorBadPoints:  &req.Review.CreatorBadPoints,
		TasteComments:     &req.Review.TasteComments,
		CreateDate:        stringToTime(req.CreateDate),
	}

	diaryFlavorItems := []flavor.PostDiaryFlavor{}
	for _, diaryFlavor := range req.DiaryFlavorList {
		diaryFlavorID, err := uuid.New().MarshalBinary()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		d := flavor.PostDiaryFlavor{
			ID:           diaryFlavorID,
			UserFlavorID: lib.ParseUUIDStrToBin(diaryFlavor.ID),
			DiaryID:      diaryID,
			Amount:       diaryFlavor.Amount,
		}
		diaryFlavorItems = append(diaryFlavorItems, d)
	}

	userDiariesItem := user.UserDiary{
		ID:      userDiariesID,
		UserID:  lib.ParseUUIDStrToBin(req.UserID),
		DiaryID: diaryID,
	}

	err = diaryEquipments.Add(diaryEquipmentsItem)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	err = diaries.Add(diaryItem)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	err = diaryFlavors.Add(diaryFlavorItems)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	err = userDiaries.Add(userDiariesItem)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successful operation"})
}
