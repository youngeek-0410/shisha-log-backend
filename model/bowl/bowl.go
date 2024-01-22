package bowl

import (
	"shisha-log-backend/lib"

	"github.com/google/uuid"
)

type Bowl struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	BowlBrand uuid.UUID `json:"brand_id"`
}

type Bowls struct {
	Items []Bowl
}

type UserBowl struct {
	BowlID    uuid.UUID `json:"id"`
	BowlName  string    `gorm:"column:name" json:"bowl_name"`
	BrandName string    `gorm:"column:name" json:"bowl_brand"`
}

type UserBowls struct {
	Items []UserBowl
}

func New() *Bowls {
	return &Bowls{}
}

func NewUserBowls() *UserBowls {
	return &UserBowls{}
}

func (r *Bowls) GetAll() []Bowl {
	db := lib.GetDBConn().DB
	var bowls []Bowl
	if err := db.Find(&bowls).Error; err != nil {
		return nil
	}
	return bowls
}

func (r *UserBowls) UserBowls(user_id string) []UserBowl {
	db := lib.GetDBConn().DB
	var userBowls []UserBowl
	binaryUUID := lib.ParseUUIDStrToBin(user_id)

	if err := db.Table("user_bowl").Select("user_bowl.bowl_id, bowl.name, bowl_brand.name").Joins("inner join bowl on user_bowl.bowl_id = bowl.id").Joins("inner join bowl_brand on bowl.brand_id = bowl_brand.id").Where("user_bowl.user_id = ?", binaryUUID).Find(&userBowls).Error; err != nil {
		return nil
	}

	return userBowls
}

// func (r *Bottles) Add(d model.Bottle) {
// 	r.Items = append(r.Items, d)
// 	db := lib.GetDBConn().DB
// 	if err := db.Create(d).Error; err != nil {
// 		fmt.Println("err!")
// 	}
// }
