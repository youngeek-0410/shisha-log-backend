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
	"time"

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
	var imagePath string

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	diaryID := uuid.New().String()
	diaryEquipmentsID := uuid.New().String()
	userDiariesID := uuid.New().String()
	imageID := uuid.New().String()

	if req.Image != "" {
		fileIdentifier := time.Now().Format("20060102-150405")
		fileName := fileIdentifier + ".jpg"
		dir := os.Getenv("IMAGE_STORAGE_PATH") + req.UserID
		imagePath = filepath.Join(dir + "/" + fileName)
		base64ImageData := strings.Split(req.Image, ",")[1]
		imageData, err := base64.StdEncoding.DecodeString(base64ImageData)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Base64データのデコードに失敗しました"})
			return
		}

		if os.Getenv("API_REVISION") == "release" {
			imageFile, _ := os.Create(fileName)
			defer imageFile.Close()

			_, err = imageFile.Write(imageData)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "ファイルの作成に失敗しました"})
				return
			}

			if _, err = imageFile.Seek(0, 0); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "ファイルポインタのリセットに失敗しました"})
				return
			}

			if err = lib.UploadFile(os.Getenv("GCS_BUCKET"), imagePath, imageFile); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "ファイルのアップロードに失敗しました"})
				return
			}

			os.Remove(fileName)
		} else {
			// ファイルを保存するディレクトリが存在していなければ作成する
			if _, err := os.Stat(dir); os.IsNotExist(err) {
				os.MkdirAll(dir, 0755)
			}

			if err := os.WriteFile(imagePath, imageData, 0666); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "ファイルの書き込みに失敗しました"})
				return
			}
		}
	} else {
		imagePath = ""
	}

	diaryImageItem := image.DiaryImage{
		ID:   imageID,
		Path: os.Getenv("GCS_BUCKET_PATH") + imagePath,
	}

	if err := diaryImages.Add(diaryImageItem); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	diaryEquipmentsItem := equipment.DiaryEquipment{
		ID:                   diaryEquipmentsID,
		UserBowlID:           req.Equipments.BowlID,
		UserBottleID:         req.Equipments.BottleID,
		UserHeatManagementID: req.Equipments.HeatManagementID,
		UserCharcoalID:       req.Equipments.CharcoalID,
		DiaryImageID:         imageID,
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
		CreateDate:        req.CreateDate,
	}

	diaryFlavorItems := []flavor.PostDiaryFlavor{}
	for _, diaryFlavor := range req.DiaryFlavorList {
		diaryFlavorID := uuid.New().String()

		d := flavor.PostDiaryFlavor{
			ID:           diaryFlavorID,
			UserFlavorID: diaryFlavor.ID,
			DiaryID:      diaryID,
			Amount:       diaryFlavor.Amount,
		}
		diaryFlavorItems = append(diaryFlavorItems, d)
	}

	userDiariesItem := user.UserDiary{
		ID:      userDiariesID,
		UserID:  req.UserID,
		DiaryID: diaryID,
	}

	if err := diaryEquipments.Add(diaryEquipmentsItem); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := diaries.Add(diaryItem); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := diaryFlavors.Add(diaryFlavorItems); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := userDiaries.Add(userDiariesItem); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successful operation"})
}
