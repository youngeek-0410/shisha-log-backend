package diary

import (
	"shisha-log-backend/lib"
	"shisha-log-backend/model/flavor"
	"time"

	"github.com/google/uuid"
)

type UserDiary struct {
	ID                uuid.UUID            `gorm:"column:id" json:"id"`
	DiaryFlavors      []flavor.DiaryFlavor `gorm:"foreignKey:ID" json:"diary_flavor_list"`
	CreateDate        time.Time            `gorm:"column:create_date" json:"create_date"`
	CreatorEvaluation string               `gorm:"column:creator_evaluation" json:"creator_evaluation"`
	TasteEvaluation   string               `gorm:"column:taste_evaluation" json:"taste_evaluation"`
}

type UserDiaries struct {
	Items []UserDiary `json:"user_diary_list"`
}

func NewUserDiaries() *UserDiaries {
	return &UserDiaries{}
}

func (r *UserDiaries) UserDiaries(user_id string) ([]UserDiary, error) {
	db := lib.GetDBConn().DB
	var userDiaries []UserDiary
	userUUID := lib.ParseUUIDStrToBin(user_id)

	if err := db.Table("diaries").Select("diaries.id, diaries.create_date, diaries.creator_evaluation, diaries.taste_evaluation").Joins("inner join user_diaries on user_diaries.diary_id = diaries.id").Where("user_diaries.user_id = ?", userUUID).Find(&userDiaries).Error; err != nil {
		return nil, err
	}

	for i := range userDiaries {
		var diaryFlavors flavor.DiaryFlavors
		diaryStrUUID := userDiaries[i].ID.String()

		flavors, err := diaryFlavors.DiaryFlavors(diaryStrUUID)
		if err != nil {
			return nil, err
		}

		userDiaries[i].DiaryFlavors = flavors
	}

	return userDiaries, nil
}
