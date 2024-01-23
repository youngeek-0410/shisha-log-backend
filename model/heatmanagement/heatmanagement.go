package heatmanagement

import (
	"shisha-log-backend/lib"

	"github.com/google/uuid"
)

type HeatManagement struct {
	ID                  uuid.UUID `json:"id"`
	Name                string    `json:"name"`
	HeatManagementBrand uuid.UUID `json:"brand_id"`
}

type HeatManagements struct {
	Items []HeatManagement
}

type UserHeatManagement struct {
	HeatManagementID   uuid.UUID `json:"id"`
	HeatManagementName string    `gorm:"column:name" json:"heat_management_name"`
	BrandName          string    `gorm:"column:name" json:"brand_name"`
}

type UserHeatManagements struct {
	Items []UserHeatManagement
}

func New() *HeatManagements {
	return &HeatManagements{}
}

func NewUserHeatManagements() *UserHeatManagements {
	return &UserHeatManagements{}
}

func (r *UserHeatManagements) UserHeatManagements(user_id string) []UserHeatManagement {
	db := lib.GetDBConn().DB
	var userHeatManagements []UserHeatManagement
	binaryUUID := lib.ParseUUIDStrToBin(user_id)

	if err := db.Table("user_heat_management").Select("user_heat_management.heat_management_id, heat_management.name, heat_management_brand.name").Joins("inner join heat_management on user_heat_management.heat_management_id = heat_management.id").Joins("inner join heat_management_brand on heat_management.brand_id = heat_management_brand.id").Where("user_heat_management.user_id = ?", binaryUUID).Find(&userHeatManagements).Error; err != nil {
		return nil
	}

	return userHeatManagements
}
