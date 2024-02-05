package diary

import (
	"shisha-log-backend/lib"
	"time"

	"github.com/google/uuid"
)

type UserDiary struct {
	ID uuid.UUID `gorm:"column:id" json:"id"`
	// DiaryFlavors      []flavor.DiaryFlavor `json:"diary_flavor_list"`
	CreateDate        time.Time `gorm:"column:create_date" json:"create_date"`
	CreatorEvaluation string    `gorm:"column:creator_evaluation" json:"creator_evaluation"`
	TasteEvaluation   string    `gorm:"column:taste_evaluation" json:"taste_evaluation"`
}

type Test struct {
	ID                uuid.UUID `gorm:"column:id" json:"id"`
	CreateDate        time.Time `gorm:"column:create_date" json:"create_date"`
	CreatorEvaluation string    `gorm:"column:creator_evaluation" json:"creator_evaluation"`
	TasteEvaluation   string    `gorm:"column:taste_evaluation" json:"taste_evaluation"`
}

// type Test struct {
// 	DiaryFlavors []flavor.DiaryFlavor `json:"diary_flavor_list"`
// }

type UserDiaries struct {
	Items []UserDiary `json:"user_diary_list"`
}

// type Diary struct {
// 	ID                uuid.UUID
// 	DiaryEquipmentsID uuid.UUID
// 	SuckingText       string
// 	CreatorEvaluation float64
// 	TasteEvaluation   float64
// 	CreatorGoodPoints string
// 	CreatorBadPoints  string
// 	TasteComments     string
// 	CreateDate        time.Time
// }

func NewUserDiaries() *UserDiaries {
	return &UserDiaries{}
}

// func (r *Diaries) Add(d model.Diary) {
// 	r.Items = append(r.Items, d)
// 	db := lib.GetDBConn().DB
// 	if err := db.Create(d).Error; err != nil {
// 		fmt.Println("err!")
// 	}
// }

func (r *UserDiaries) UserDiaries(user_id string) ([]UserDiary, error) {
	db := lib.GetDBConn().DB
	var userDiaries []UserDiary
	// var test []Test
	userUUID := lib.ParseUUIDStrToBin(user_id)

	// if err := db.Table("diaries").Joins("inner join user_diaries on user_diaries.diary_id = diaries.id").Where("user_diaries.user_id = ?", userUUID).Find(&test).Error; err != nil {
	if err := db.Table("diaries").Joins("inner join user_diaries on user_diaries.diary_id = diaries.id").Where("user_diaries.user_id = ?", userUUID).Find(&userDiaries).Error; err != nil {
		return nil, err
	}

	// for _, diary := range userDiaries {
	// 	var diaryFlavors flavor.DiaryFlavors
	// 	diaryStrUUID := diary.ID.String()

	// 	flavors, err := diaryFlavors.DiaryFlavors(diaryStrUUID)
	// 	if err != nil {
	// 		return nil, err
	// 	}

	// 	diary.DiaryFlavors = flavors
	// }

	return userDiaries, nil
}

// func (r *UserDiaries) UserDiaries(userID string) ([]UserDiary, error) {
// var userDiaries []UserDiary
// db := lib.GetDBConn().DB
// binaryUUID := lib.ParseUUIDStrToBin(userID)

// ユーザーの日記を取得
// var userDiaries []UserDiary
// err := db.Table("user_diaries").Select("user_diaries.ID, diaries.sucking_text").Joins("inner join diaries on user_diaries.diary_id = diaries.id").Where("user_id = ?", binaryUUID).Find(&userDiaries).Error
// if err != nil {
// 	return nil, err
// }

// 各日記に対して、フレーバー情報を取得
// for _, diary := range userDiaries {
// 	var diaryFlavors []DiaryFlavor
// 	for _, df := range diary.DiaryFlavorList {
// 		diaryFlavors = append(diaryFlavors, DiaryFlavor{
// 			ID: df.ID,
// FlavorName: df.FlavorName,
// BrandName:  df.BrandName,
// 			Amount: float64(df.Amount),
// 		})
// 	}

// 	userDiaries = append(userDiaries, UserDiary{
// 		ID:                diary.ID,
// 		DiaryFlavorList:   diaryFlavors,
// 		CreateDate:        diary.CreateDate,
// 		CreatorEvaluation: diary.CreatorEvaluation,
// 		TasteEvaluation:   diary.TasteEvaluation,
// 	})
// }

// 	return userDiaries, nil
// }
