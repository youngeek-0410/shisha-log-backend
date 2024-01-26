package charcoal

import (
	"shisha-log-backend/lib"

	"github.com/google/uuid"
)

// type Charcoal struct {
// 	ID            uuid.UUID `json:"id"`
// 	Name          string    `json:"name"`
// 	CharcoalBrand uuid.UUID `json:"brand_id"`
// }

// type Charcoals struct {
// 	Items []Charcoal
// }

type UserCharcoal struct {
	CharcoalID   uuid.UUID `json:"id"`
	CharcoalName string    `gorm:"column:name" json:"charcoal_name"`
	BrandName    string    `gorm:"column:name" json:"brand_name"`
}

type UserCharcoals struct {
	Items []UserCharcoal
}

// func New() *Charcoals {
// 	return &Charcoals{}
// }

func NewUserCharcoals() *UserCharcoals {
	return &UserCharcoals{}
}

func (r *UserCharcoals) UserCharcoals(user_id string) ([]UserCharcoal, error) {
	db := lib.GetDBConn().DB
	var userCharcoals []UserCharcoal
	binaryUUID := lib.ParseUUIDStrToBin(user_id)

	if err := db.Table("user_charcoals").Select("user_charcoals.charcoal_id, charcoals.name, charcoal_brands.name").Joins("inner join charcoals on user_charcoals.charcoal_id = charcoals.id").Joins("inner join charcoal_brands on charcoals.brand_id = charcoal_brands.id").Where("user_charcoals.user_id = ?", binaryUUID).Find(&userCharcoals).Error; err != nil {
		return nil, err
	}

	return userCharcoals, nil
}
