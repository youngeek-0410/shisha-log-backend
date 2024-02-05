package bowl

import (
	"shisha-log-backend/lib"

	"github.com/google/uuid"
)

type UserBowl struct {
	BowlID    uuid.UUID `json:"id"`
	BowlName  string    `gorm:"column:name" json:"bowl_name"`
	BrandName string    `gorm:"column:name" json:"brand_name"`
}

type UserBowls struct {
	Items []UserBowl
}

func NewUserBowls() *UserBowls {
	return &UserBowls{}
}

func (r *UserBowls) UserBowls(user_id string) ([]UserBowl, error) {
	db := lib.GetDBConn().DB
	var userBowls []UserBowl
	binaryUUID := lib.ParseUUIDStrToBin(user_id)

	if err := db.Table("user_bowls").Select("user_bowls.bowl_id, bowls.name, bowl_brands.name").Joins("inner join bowls on user_bowls.bowl_id = bowls.id").Joins("inner join bowl_brands on bowls.brand_id = bowl_brands.id").Where("user_bowls.user_id = ?", binaryUUID).Find(&userBowls).Error; err != nil {
		return nil, err
	}

	return userBowls, nil
}

// func (r *Bottles) Add(d model.Bottle) {
// 	r.Items = append(r.Items, d)
// 	db := lib.GetDBConn().DB
// 	if err := db.Create(d).Error; err != nil {
// 		fmt.Println("err!")
// 	}
// }
