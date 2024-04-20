package equipment

import (
	"shisha-log-backend/lib"
	"time"
)

type DiaryEquipment struct {
	ID                   string     `gorm:"column:id"`
	UserBowlID           string     `gorm:"column:user_bowl_id"`
	UserBottleID         string     `gorm:"column:user_bottle_id"`
	UserHeatManagementID string     `gorm:"column:user_heat_management_id"`
	UserCharcoalID       string     `gorm:"column:user_charcoal_id"`
	DiaryImageID         string     `gorm:"column:diary_image_id"`
	CreatedAt            *time.Time `gorm:"column:created_at"`
	UpdatedAt            *time.Time `gorm:"column:updated_at"`
}

type DiaryEquipments struct {
	Items []DiaryEquipment
}

func (r *DiaryEquipments) Add(d DiaryEquipment) error {
	r.Items = append(r.Items, d)
	db := lib.GetDBConn().DB
	if err := db.Create(&d).Error; err != nil {
		return err
	}
	return nil
}
