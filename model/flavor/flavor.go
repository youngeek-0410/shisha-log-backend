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
	FlavorID   uuid.UUID `json:"id"`
	FlavorName string    `gorm:"column:name" json:"flavor_name"`
	BrandName  string    `gorm:"column:name" json:"brand_name"`
}

type UserFlavors struct {
	Items []UserFlavor
}

type DiaryFlavor struct {
	ID         uuid.UUID
	FlavorName string
	BrandName  string
	Amount     float64
}

type DiaryFlavors struct {
	Items []DiaryFlavor
}

func New() *Flavors {
	return &Flavors{}
}

func NewUserFlavors() *UserFlavors {
	return &UserFlavors{}
}

func NewDiaryFlavors() *DiaryFlavors {
	return &DiaryFlavors{}
}

func (r *UserFlavors) UserFlavors(user_id string) ([]UserFlavor, error) {
	db := lib.GetDBConn().DB
	var userFlavors []UserFlavor
	binaryUUID := lib.ParseUUIDStrToBin(user_id)

	if err := db.Table("user_flavors").Select("user_flavors.flavor_id, flavors.name, flavor_brands.name").Joins("inner join flavors on user_flavors.flavor_id = flavors.id").Joins("inner join flavor_brands on flavors.brand_id = flavor_brands.id").Where("user_flavors.user_id = ?", binaryUUID).Find(&userFlavors).Error; err != nil {
		return nil, err
	}

	return userFlavors, nil
}

func (r *DiaryFlavors) DiaryFlavors(diary_id string) ([]DiaryFlavor, error) {
	db := lib.GetDBConn().DB
	var diaryFlavors []DiaryFlavor
	diaryUUID := lib.ParseUUIDStrToBin(diary_id)

	if err := db.Table("diary_flavors").Select("flavors.id, flavors.name, flavor_brands.name, diary_flavors.amount").Joins("inner join flavors on diary_flavors.user_flavor_id = flavors.id").Joins("inner join flavor_brands on flavors.brand_id = flavor_brands.id").Where("diary_id = ?", diaryUUID).Find(&diaryFlavors).Error; err != nil {
		return nil, err
	}

	return diaryFlavors, nil
}
