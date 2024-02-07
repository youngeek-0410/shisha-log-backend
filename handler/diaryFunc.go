package handler

import (
	"net/http"
	"shisha-log-backend/model/diary"

	"github.com/gin-gonic/gin"
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
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// db :=

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
