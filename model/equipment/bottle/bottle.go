package bottle

import (
	"shisha-log-backend/lib"
)

type UserBottle struct {
	BottleID   string `json:"id"`
	BottleName string `gorm:"column:name" json:"bottle_name"`
	BrandName  string `gorm:"column:name" json:"brand_name"`
}

type UserBottles struct {
	Items []UserBottle
}

func NewUserBottles() *UserBottles {
	return &UserBottles{}
}

func (r *UserBottles) UserBottles(user_id string) ([]UserBottle, error) {
	db := lib.GetDBConn().DB
	var userBottles []UserBottle

	if err := db.Table("user_bottles").Select("user_bottles.bottle_id, bottles.name, bottle_brands.name").Joins("inner join bottles on user_bottles.bottle_id = bottles.id").Joins("inner join bottle_brands on bottles.brand_id = bottle_brands.id").Where("user_bottles.user_id = ?", user_id).Find(&userBottles).Error; err != nil {
		return nil, err
	}

	return userBottles, nil
}

// func (r *Bottles) Add(d model.Bottle) {
// 	r.Items = append(r.Items, d)
// 	db := lib.GetDBConn().DB
// 	if err := db.Create(d).Error; err != nil {
// 		fmt.Println("err!")
// 	}
// }
