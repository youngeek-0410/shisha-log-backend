package flavor

import (
	"shisha-log-backend/lib"
	"time"
)

type GetUserFlavorResponseItem struct {
	FlavorID   string `json:"id"`
	FlavorName string `gorm:"column:name" json:"flavor_name"`
	BrandName  string `gorm:"column:name" json:"brand_name"`
}

type GetUserFlavorResponse struct {
	Items []GetUserFlavorResponseItem
}

type DiaryFlavor struct {
	ID         string  `json:"id"`
	FlavorName string  `gorm:"column:name" json:"flavor_name"`
	BrandName  string  `gorm:"column:name" json:"brand_name"`
	Amount     float64 `json:"amount"`
}

type DiaryFlavors struct {
	Items []DiaryFlavor
}

type PostDiaryFlavor struct {
	ID           string     `gorm:"column:id"`
	UserFlavorID string     `gorm:"column:user_flavor_id"`
	DiaryID      string     `gorm:"column:diary_id"`
	Amount       float64    `gorm:"column:amount"`
	CreatedAt    *time.Time `gorm:"column:created_at"`
	UpdatedAt    *time.Time `gorm:"column:updated_at"`
}

type PostDiaryFlavors struct {
	Items []PostDiaryFlavor
}

type FlavorBrand struct {
	ID        string
	Name      string
	CreatedAt *time.Time
}

type Flavor struct {
	ID         string  `gorm:"id"`
	BrandID    string  `gorm:"brand_id"`
	Name       string  `gorm:"name"`
	CreateArea *string `gorm:"name"`
	CreatedAt  *time.Time
}

type UserFlavor struct {
	ID        string
	FlavorID  string
	UserID    string
	CreatedAt *time.Time
}

type PostFlavorRequest struct {
	UserID     string  `json:"user_id"`
	BrandID    string  `json:"brand_id"`
	Name       string  `json:"name"`
	CreateArea *string `json:"create_area"`
}

func NewUserFlavorResponse() *GetUserFlavorResponse {
	return &GetUserFlavorResponse{}
}

func NewDiaryFlavors() *DiaryFlavors {
	return &DiaryFlavors{}
}

func (r *GetUserFlavorResponse) GetUserFlavorResponse(user_id string) ([]GetUserFlavorResponseItem, error) {
	db := lib.GetDBConn().DB
	var GetUserFlavorResponse []GetUserFlavorResponseItem

	if err := db.Table("user_flavors").Select("user_flavors.flavor_id, flavors.name, flavor_brands.name").Joins("inner join flavors on user_flavors.flavor_id = flavors.id").Joins("inner join flavor_brands on flavors.brand_id = flavor_brands.id").Where("user_flavors.user_id = ?", user_id).Find(&GetUserFlavorResponse).Error; err != nil {
		return nil, err
	}

	return GetUserFlavorResponse, nil
}

func (r *DiaryFlavors) DiaryFlavors(diary_id string) ([]DiaryFlavor, error) {
	db := lib.GetDBConn().DB
	var diaryFlavors []DiaryFlavor

	if err := db.Table("diary_flavors").Select("flavors.id, flavors.name, flavor_brands.name, diary_flavors.amount").Joins("inner join flavors on diary_flavors.user_flavor_id = flavors.id").Joins("inner join flavor_brands on flavors.brand_id = flavor_brands.id").Where("diary_id = ?", diary_id).Find(&diaryFlavors).Error; err != nil {
		return nil, err
	}

	return diaryFlavors, nil
}

func (r *PostDiaryFlavors) Add(d []PostDiaryFlavor) error {
	r.Items = append(r.Items, d...)
	db := lib.GetDBConn().DB
	if err := db.Table("diary_flavors").Create(&d).Error; err != nil {
		return err
	}
	return nil
}

func (r *FlavorBrand) Add(flavorBrand FlavorBrand) error {
	db := lib.GetDBConn().DB
	if err := db.Create(&flavorBrand).Error; err != nil {
		return err
	}
	return nil
}

func (r *Flavor) Add(flavor Flavor) error {
	db := lib.GetDBConn().DB
	if err := db.Create(&flavor).Error; err != nil {
		return err
	}
	return nil
}

func (r *UserFlavor) Add(userFlavor UserFlavor) error {
	db := lib.GetDBConn().DB
	if err := db.Create(&userFlavor).Error; err != nil {
		return err
	}
	return nil
}
