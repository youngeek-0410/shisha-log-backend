package equipment

import (
	"fmt"
	"shisha-log-backend/lib"
	"time"
)

type DiaryEquipment struct {
	ID                   []byte     `gorm:"column:id"`
	UserBowlID           []byte     `gorm:"column:user_bowl_id"`
	UserBottleID         []byte     `gorm:"column:user_bottle_id"`
	UserHeatManagementID []byte     `gorm:"column:user_heat_management_id"`
	UserCharcoalID       []byte     `gorm:"column:user_charcoal_id"`
	DiaryImageID         []byte     `gorm:"column:diary_image_id"`
	CreatedAt            *time.Time `gorm:"column:created_at"`
	UpdatedAt            *time.Time `gorm:"column:updated_at"`
}

type DiaryEquipments struct {
	Items []DiaryEquipment
}

func (r *DiaryEquipments) DiaryEquipmentsAdd(d DiaryEquipment) {
	r.Items = append(r.Items, d)
	db := lib.GetDBConn().DB
	if err := db.Create(&d).Error; err != nil {
		fmt.Println("err!")
	}
}
