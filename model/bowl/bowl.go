package bowl

import (
	"shisha-log-backend/lib"
	"shisha-log-backend/model"

	"github.com/google/uuid"
)

type Bowls struct {
	Items []model.Bowl
}

type UserBowls struct {
	Items []UserBowl
}

type UserBowl struct {
	BowlID    uuid.UUID `json:"id"`
	BowlName  string    `gorm:"column:name" json:"bowl_name"`
	BrandName string    `gorm:"column:name" json:"bowl_brand"`
}

func New() *Bowls {
	return &Bowls{}
}

// func (r *Bottles) Add(d model.Bottle) {
// 	r.Items = append(r.Items, d)
// 	db := lib.GetDBConn().DB
// 	if err := db.Create(d).Error; err != nil {
// 		fmt.Println("err!")
// 	}
// }

func (r *Bowls) GetAll() []model.Bowl {
	db := lib.GetDBConn().DB
	var bowls []model.Bowl
	if err := db.Find(&bowls).Error; err != nil {
		return nil
	}
	return bowls
}

func NewUserBowls() *UserBowls {
	return &UserBowls{}
}

func (r *UserBowls) UserBowls(user_id string) []UserBowl {
	db := lib.GetDBConn().DB
	var userBowls []UserBowl

	userID, parseErr := uuid.Parse(user_id)
	if parseErr != nil {
		return nil
	}
	binaryUUID, err := userID.MarshalBinary()
	if err != nil {
		return nil
	}
	if err := db.Table("user_bowl").Select("user_bowl.bowl_id, bowl.name, bowl_brand.name").Joins("inner join bowl on user_bowl.bowl_id = bowl.id").Joins("inner join bowl_brand on bowl.brand_id = bowl_brand.id").Where("user_bowl.user_id = ?", binaryUUID).Find(&userBowls).Error; err != nil {
		return nil
	}

	return userBowls
}
