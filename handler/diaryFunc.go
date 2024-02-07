package handler

import (
	"fmt"
	"net/http"
	"shisha-log-backend/model/diary"

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
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	diaryID := uuid.New()
	diaryEquipmentsID := uuid.New()

	item := diary.Diary{
		ID:                diaryID,
		DiaryEquipmentsID: diaryEquipmentsID,
		// SuckingText:       &req.SuckingText,
		// Temperature:       &req.Equipments.Climate.Temperature,
		// Humidity:          &req.Equipments.Climate.Humidity,
		CreatorEvaluation: req.Review.CreatorEvaluation,
		TasteEvaluation:   req.Review.TasteEvaluation,
		// CreatorGoodPoints: &req.Review.CreatorGoodPoints,
		// CreatorBadPoints:  &req.Review.CreatorBadPoints,
		// TasteComments:     &req.Review.TasteComments,
		CreateDate: stringToTime(req.CreateDate),
	}

	fmt.Println(item)
	diaries.Add(item)

	c.JSON(http.StatusOK, gin.H{"message": "Successful operation"})
}

// type DiaryPostRequest struct {
// 	Id    int    `json:"id"`
// 	Title string `json:"title"`
// }

// func DiaryPost(post *diary.Diaries) gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		requestBody := DiaryPostRequest{}
// 		c.Bind(&requestBody)

// 		item := diary.Diary{
// 			Id:    requestBody.Id,
// 			Title: requestBody.Title,
// 		}
// 		post.Add(item)

// 		c.Status(http.StatusNoContent)
// 	}
// }
