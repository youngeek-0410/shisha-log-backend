package handler

import (
	"encoding/base64"
	"net/http"
	"os"
	"path/filepath"
	"shisha-log-backend/lib"
	"shisha-log-backend/model/diary"
	"shisha-log-backend/model/equipment"
	"shisha-log-backend/model/flavor"
	"shisha-log-backend/model/image"
	"shisha-log-backend/model/user"
	"strings"

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
		c.JSON(http.StatusOK, gin.H{
			"user_diary_list": result,
		})
	}
}

func CreateDiary(c *gin.Context) {
	var req diary.DiaryRequest
	var diaries diary.Diaries
	var diaryEquipments equipment.DiaryEquipments
	var diaryFlavors flavor.PostDiaryFlavors
	var userDiaries user.UserDiaries
	var diaryImages image.DiaryImages
	var imageFullPath string

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

	imageID := uuid.New()
	imageBinaryID, err := imageID.MarshalBinary()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if req.Image != "" {
		// ファイルを保存するディレクトリが存在していなければ作成する
		dir := os.Getenv("IMAGE_STORAGE_PATH") + req.UserID
		if _, err := os.Stat(dir); os.IsNotExist(err) {
			os.MkdirAll(dir, 0755)
		}

		imageFullPath = filepath.Join(dir + "/" + imageID.String() + ".jpg")
		base64ImageData := strings.Split(req.Image, ",")[1]
		imageData, err := base64.StdEncoding.DecodeString(base64ImageData)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Base64データのデコードに失敗しました"})
			return
		}

		if err := os.WriteFile(imageFullPath, imageData, 0666); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "ファイルの書き込みに失敗しました"})
			return
		}
	} else {
		imageFullPath = ""
	}

	diaryImageItem := image.DiaryImage{
		ID:   imageBinaryID,
		Path: imageFullPath,
	}

	err = diaryImages.Add(diaryImageItem)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	diaryEquipmentsItem := equipment.DiaryEquipment{
		ID:                   diaryEquipmentsID,
		UserBowlID:           lib.ParseUUIDStrToBin(req.Equipments.BowlID),
		UserBottleID:         lib.ParseUUIDStrToBin(req.Equipments.BottleID),
		UserHeatManagementID: lib.ParseUUIDStrToBin(req.Equipments.HeatManagementID),
		UserCharcoalID:       lib.ParseUUIDStrToBin(req.Equipments.CharcoalID),
		DiaryImageID:         imageBinaryID,
	}

	diaryItem := diary.Diary{
		ID:                diaryID,
		DiaryEquipmentsID: diaryEquipmentsID,
		ServeText:         &req.ServeText,
		SuckingText:       &req.SuckingText,
		Temperature:       &req.Equipments.Climate.Temperature,
		Humidity:          &req.Equipments.Climate.Humidity,
		CreatorEvaluation: req.Review.CreatorEvaluation,
		TasteEvaluation:   req.Review.TasteEvaluation,
		CreatorGoodPoints: &req.Review.CreatorGoodPoints,
		CreatorBadPoints:  &req.Review.CreatorBadPoints,
		TasteComments:     &req.Review.TasteComments,
		CreateDate:        lib.StringToTime(req.CreateDate),
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
