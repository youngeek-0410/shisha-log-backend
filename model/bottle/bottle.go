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
	BowlID    uuid.UUID `json:"id"`
	BowlName  string    `gorm:"column:name" json:"bowl_name"`
	BrandName string    `gorm:"column:name" json:"bowl_brand"`
}

type UserBottles struct {
	Items []UserBottle
}

func New() *Bottles {
	return &Bottles{}
}

// func (r *Bottles) Add(d model.Bottle) {
// 	r.Items = append(r.Items, d)
// 	db := lib.GetDBConn().DB
// 	if err := db.Create(d).Error; err != nil {
// 		fmt.Println("err!")
// 	}
// }

func (r *Bottles) GetAll() []Bottle {
	db := lib.GetDBConn().DB
	var bottles []Bottle
	if err := db.Find(&bottles).Error; err != nil {
		return nil
	}
	return bottles
}

func (r *UserBottles) UserBowls(user_id string) []UserBottle {
	db := lib.GetDBConn().DB
	var userBowls []UserBottle

	userID, parseErr := uuid.Parse(user_id)
	if parseErr != nil {
		return nil
	}
	binaryUUID, err := userID.MarshalBinary()
	if err != nil {
		return nil
	}
	if err := db.Table("user_bottle").Select("user_bowl.bowl_id, bowl.name, bowl_brand.name").Joins("inner join bowl on user_bowl.bowl_id = bowl.id").Joins("inner join bowl_brand on bowl.brand_id = bowl_brand.id").Where("user_bowl.user_id = ?", binaryUUID).Find(&userBowls).Error; err != nil {
		return nil
	}

	return userBowls
}
