package charcoal

import (
	"shisha-log-backend/lib"
)

type UserCharcoal struct {
	CharcoalID   string `json:"id"`
	CharcoalName string `gorm:"column:name" json:"charcoal_name"`
	BrandName    string `gorm:"column:name" json:"brand_name"`
}

type UserCharcoals struct {
	Items []UserCharcoal
}

func NewUserCharcoals() *UserCharcoals {
	return &UserCharcoals{}
}

func (r *UserCharcoals) UserCharcoals(user_id string) ([]UserCharcoal, error) {
	db := lib.GetDBConn().DB
	var userCharcoals []UserCharcoal

	if err := db.Table("user_charcoals").Select("user_charcoals.charcoal_id, charcoals.name, charcoal_brands.name").Joins("inner join charcoals on user_charcoals.charcoal_id = charcoals.id").Joins("inner join charcoal_brands on charcoals.brand_id = charcoal_brands.id").Where("user_charcoals.user_id = ?", user_id).Find(&userCharcoals).Error; err != nil {
		return nil, err
	}

	return userCharcoals, nil
}
