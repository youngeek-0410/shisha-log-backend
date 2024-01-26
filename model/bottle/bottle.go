package bottle

import (
	"shisha-log-backend/lib"

	"github.com/google/uuid"
)

type Bottle struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	BottleBrand uuid.UUID `json:"brand_id"`
}

type Bottles struct {
	Items []Bottle
}

type UserBottle struct {
	BottleID   uuid.UUID `json:"id"`
	BottleName string    `gorm:"column:name" json:"bottle_name"`
	BrandName  string    `gorm:"column:name" json:"brand_name"`
}

type UserBottles struct {
	Items []UserBottle
}

func New() *Bottles {
	return &Bottles{}
}

func NewUserBottles() *UserBottles {
	return &UserBottles{}
}

// func (r *Bottles) GetAll() []Bottle {
// 	db := lib.GetDBConn().DB
// 	var bottles []Bottle
// 	if err := db.Find(&bottles).Error; err != nil {
// 		return nil
// 	}
// 	return bottles
// }

func (r *UserBottles) UserBottles(user_id string) ([]UserBottle, error) {
	db := lib.GetDBConn().DB
	var userBottles []UserBottle
	binaryUUID := lib.ParseUUIDStrToBin(user_id)

	if err := db.Table("user_bottles").Select("user_bottles.bottle_id, bottles.name, bottle_brands.name").Joins("inner join bottles on user_bottles.bottle_id = bottles.id").Joins("inner join bottle_brands on bottles.brand_id = bottle_brands.id").Where("user_bottles.user_id = ?", binaryUUID).Find(&userBottles).Error; err != nil {
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
