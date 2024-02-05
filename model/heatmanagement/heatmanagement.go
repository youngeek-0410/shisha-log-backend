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

func NewUserHeatManagements() *UserHeatManagements {
	return &UserHeatManagements{}
}

func (r *UserHeatManagements) UserHeatManagements(user_id string) ([]UserHeatManagement, error) {
	db := lib.GetDBConn().DB
	var userHeatManagements []UserHeatManagement
	binaryUUID := lib.ParseUUIDStrToBin(user_id)

	if err := db.Table("user_heat_managements").Select("user_heat_managements.heat_management_id, heat_managements.name, heat_management_brands.name").Joins("inner join heat_managements on user_heat_managements.heat_management_id = heat_managements.id").Joins("inner join heat_management_brands on heat_managements.brand_id = heat_management_brands.id").Where("user_heat_managements.user_id = ?", binaryUUID).Find(&userHeatManagements).Error; err != nil {
		return nil, err
	}

	return userHeatManagements, nil
}
