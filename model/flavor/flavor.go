package flavor

import (
	"shisha-log-backend/lib"

	"github.com/google/uuid"
)

type Flavor struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	FlavorBrand uuid.UUID `json:"brand_id"`
}

type Flavors struct {
	Items []Flavor
}

type UserFlavor struct {
	FlavorID   uuid.UUID `json:"flavor_id"`
	FlavorName string    `gorm:"column:name" json:"flavor_name"`
	BrandName  string    `gorm:"column:name" json:"brand_name"`
}

type UserFlavors struct {
	Items []UserFlavor
}

func New() *Flavors {
	return &Flavors{}
}

func NewUserFlavors() *UserFlavors {
	return &UserFlavors{}
}

func (r *UserFlavors) UserFlavors(user_id string) []UserFlavor {
	db := lib.GetDBConn().DB
	var userFlavors []UserFlavor
	binaryUUID := lib.ParseUUIDStrToBin(user_id)

	if err := db.Table("user_flavor").Select("user_flavor.flavor_id, flavor.name, flavor_brand.name").Joins("inner join flavor on user_flavor.flavor_id = flavor.id").Joins("inner join flavor_brand on flavor.brand_id = flavor_brand.id").Where("user_flavor.user_id = ?", binaryUUID).Find(&userFlavors).Error; err != nil {
		return nil
	}

	return userFlavors
}
