package diary

import (
	"shisha-log-backend/lib"
	"time"

	"github.com/google/uuid"
)

type Diary struct {
	// Flavor           []Flavor
	// Temperature      *float64
	// Humidity         *float64
	// CreatorGoodPoint *string
	// CreatorBadPoint  *string
	// CreatorStar      int
	// TasteStar        int

	ID              uuid.UUID
	DiaryEquipments uuid.UUID
	// SuckingText       string
	CreatorEvaluation float64
	TasteEvaluation   float64
	// CreatorGoodPoints string
	// CreatorBadPoints  string
	// TasteComments     string
	// CreateDate        time.Time
	// CreatedAt         time.Time
	// UpdatedAt         time.Time
}

type DiaryEquipments struct {
	ID           uuid.UUID
	DiaryFlavors []DiaryFlavor
}

type DiaryFlavor struct {
	DiaryID    uuid.UUID `gorm:foreign`
	FlavorID   uuid.UUID `json:"flavor_id"`
	FlavorName string    `json:"flavor_name"`
	BrandName  string    `json:"brand_name"`
	Amount     float64   `json:"amount"`
}

type UserDiary struct {
	DiaryFlavorList   []DiaryFlavor `json:"user_flavor_list"`
	IsCreated         bool          `json:"is_created"`
	CreateDate        time.Time     `json:"create_date"`
	CreatorEvaluation float64       `json:"creator_evaluation"`
	TasteEvaluation   float64       `json:"taste_evaluation"`
}

type UserDiaries struct {
	Items []UserDiary `json:"user_diary_list"`
}

// type Flavor struct {
// 	ID      uuid.UUID `gorm:"type:binary(16)"`
// 	BrandID uuid.UUID `gorm:"type:binary(16)"`
// 	Name    string
// 	// ブランド情報
// 	FlavorBrand FlavorBrand `gorm:"foreignKey:BrandID;references:ID"`
// }

// type FlavorBrand struct {
// 	ID   uuid.UUID `gorm:"type:binary(16)"`
// 	Name string
// }

// func New() *Diaries {
// 	return &Diaries{}
// }

func NewUserDiary() *UserDiaries {
	return &UserDiaries{}
}

// func (r *Diaries) Add(d model.Diary) {
// 	r.Items = append(r.Items, d)
// 	db := lib.GetDBConn().DB
// 	if err := db.Create(d).Error; err != nil {
// 		fmt.Println("err!")
// 	}
// }

// func (r *UserDiaries) UserDiaries(user_id string) ([]model.Diary, error) {
// 	db := lib.GetDBConn().DB
// 	var userDiaries UserDiaries
// 	binaryUUID := lib.ParseUUIDStrToBin(user_id)

// 	if err := db.Find(&diaries).Error; err != nil {
// 		return nil, err
// 	}

// 	if err := db.Table("diary_list").Select("").Joins("inner join diary on diary_list.diary_id = dairy.id").Joins("inner join bottle_brand on bottle.brand_id = bottle_brand.id").Where("user_bottle.user_id = ?", binaryUUID).Find(&userDiaries).Error; err != nil {
// 		return nil, err
// 	}

// 	return userDiaries, nil
// }

// func (r *UserDiaries) userDiaries(userID uuid.UUID) ([]UserDiary, error) {
func (r *UserDiaries) UserDiaries(userID string) ([]UserDiary, error) {
	var userDiaries []UserDiary
	db := lib.GetDBConn().DB
	// binaryUUID := lib.ParseUUIDStrToBin(userID)

	// ユーザーの日記を取得
	var diaries []UserDiary
	// err := db.Where("user_id = ?", userID).Preload("DiaryFlavors.Flavor.FlavorBrand").Find(&diaries).Error
	// err := db.Preload("DiaryFlavors.Flavor.FlavorBrand").Where("user_id = ?", binaryUUID).Find(&diaries).Error
	err := db.Table("diary_list").Joins("inner join diary on diary_list.diary_id = diary.id").Joins("inner join diary_flavors on diary.id = diary_flavors.diary_id").Joins("inner join user_flavor on diary_flavors.user_flavor_id = user_flavors.flavor_id").Joins("inner join flavor on user_flavor.flavor_id = flavor.id").Joins("inner join flavor_brand on flavor.brand_id = flavor_brand.id").Find(&userDiaries).Error
	if err != nil {
		return nil, err
	}

	// 各日記に対して、フレーバー情報を取得
	for _, diary := range diaries {
		var diaryFlavors []DiaryFlavor
		for _, df := range diary.DiaryFlavorList {
			diaryFlavors = append(diaryFlavors, DiaryFlavor{
				FlavorID:   df.FlavorID,
				FlavorName: df.FlavorName,
				BrandName:  df.BrandName,
				Amount:     float64(df.Amount),
			})
		}

		userDiaries = append(userDiaries, UserDiary{
			DiaryFlavorList:   diaryFlavors,
			IsCreated:         diary.IsCreated,
			CreateDate:        diary.CreateDate,
			CreatorEvaluation: diary.CreatorEvaluation,
			TasteEvaluation:   diary.TasteEvaluation,
		})
	}

	return userDiaries, nil
}
