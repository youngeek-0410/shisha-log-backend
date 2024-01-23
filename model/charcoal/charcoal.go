package charcoal

import (
	"shisha-log-backend/lib"

	"github.com/google/uuid"
)

type Charcoal struct {
	ID            uuid.UUID `json:"id"`
	Name          string    `json:"name"`
	CharcoalBrand uuid.UUID `json:"brand_id"`
}

type Charcoals struct {
	Items []Charcoal
}

type UserCharcoal struct {
	CharcoalID   uuid.UUID `json:"id"`
	CharcoalName string    `gorm:"column:name" json:"charcoal_name"`
	BrandName    string    `gorm:"column:name" json:"charcoal_brand"`
}

type UserCharcoals struct {
	Items []UserCharcoal
}

func New() *Charcoals {
	return &Charcoals{}
}

func NewUserCharcoals() *UserCharcoals {
	return &UserCharcoals{}
}

func (r *UserCharcoals) UserCharcoals(user_id string) []UserCharcoal {
	db := lib.GetDBConn().DB
	var userCharcoals []UserCharcoal
	binaryUUID := lib.ParseUUIDStrToBin(user_id)

	if err := db.Table("user_charcoal").Select("user_charcoal.charcoal_id, charcoal.name, charcoal_brand.name").Joins("inner join charcoal on user_charcoal.charcoal_id = charcoal.id").Joins("inner join charcoal_brand on charcoal.brand_id = charcoal_brand.id").Where("user_charcoal.user_id = ?", binaryUUID).Find(&userCharcoals).Error; err != nil {
		return nil
	}

	return userCharcoals
}
