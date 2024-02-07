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
	CreatorEvaluation float64              `gorm:"column:creator_evaluation" json:"creator_evaluation"`
	TasteEvaluation   float64              `gorm:"column:taste_evaluation" json:"taste_evaluation"`
}

type UserDiaries struct {
	Items []UserDiary `json:"user_diary_list"`
}

type DiaryRequest struct {
	UserID          string          `json:"user_id"`
	Equipments      DiaryEquipments `json:"equipments"`
	DiaryFlavorList []FlavorAmount  `json:"diary_flavor_list"`
	ImageID         string          `json:"image_id"`
	SuckingText     string          `json:"sucking_text"`
	Review          DiaryReview     `json:"review"`
	CreateDate      string          `json:"create_date"`
}

type DiaryEquipments struct {
	BowlID           string  `json:"bowl_id"`
	BottleID         string  `json:"bottle_id"`
	HeatManagementID string  `json:"heat_management_id"`
	CharcoalID       string  `json:"charcoal_id"`
	Climate          Climate `json:"climate"`
}

type Climate struct {
	Temperature float64 `json:"temperature"`
	Humidity    float64 `json:"humidity"`
}

type FlavorAmount struct {
	ID     string  `json:"id"`
	Amount float64 `json:"amount"`
}

type DiaryReview struct {
	CreatorEvaluation float64 `json:"creator_evaluation"`
	TasteEvaluation   float64 `json:"taste_evaluation"`
	CreatorGoodPoints string  `json:"creator_good_points"`
	CreatorBadPoints  string  `json:"creator_bad_points"`
	TasteComments     string  `json:"taste_comments"`
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

// func (r *Diaries) Add(d Diary) {
// 	r.Items = append(r.Items, d)
// 	db := lib.GetDBConn().DB
// 	if err := db.Create(d).Error; err != nil {
// 		fmt.Println("err!")
// 	}
// }
