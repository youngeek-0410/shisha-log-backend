package fravor

import (
	"shisha-log-backend/lib"

	"github.com/google/uuid"
)

type Fravor struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	FravorBrand uuid.UUID `json:"brand_id"`
}

type Fravors struct {
	Items []Fravor
}

type UserFravor struct {
	FravorID   uuid.UUID `json:"fravor_id"`
	FravorName string    `gorm:"column:name" json:"fravor_name"`
	BrandName  string    `gorm:"column:name" json:"brand_name"`
}

type UserFravors struct {
	Items []UserFravor
}

func New() *Fravors {
	return &Fravors{}
}

func NewUserFravors() *UserFravors {
	return &UserFravors{}
}

func (r *UserFravors) UserFravors(user_id string) []UserFravor {
	db := lib.GetDBConn().DB
	var userFravors []UserFravor
	binaryUUID := lib.ParseUUIDStrToBin(user_id)

	if err := db.Table("user_fravor").Select("user_fravor.fravor_id, fravor.name, fravor_brand.name").Joins("inner join fravor on user_fravor.fravor_id = fravor.id").Joins("inner join fravor_brand on fravor.brand_id = fravor_brand.id").Where("user_fravor.user_id = ?", binaryUUID).Find(&userFravors).Error; err != nil {
		return nil
	}

	return userFravors
}
