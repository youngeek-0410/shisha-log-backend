package model

import (
	"time"

	"github.com/google/uuid"
)

type Base struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primry_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

type User struct {
	Base
	Name    string
	Diaries *Diaries
}

type Diary struct {
	Base
	Fravor           []Fravor
	Bottle           Bottle
	Bowl             Bowl
	HeatManagement   HeatManagement
	Chacoal          Chacoal
	Temperature      *float64
	Humidity         *float64
	CreatorGoodPoint *string
	CreatorBadPoint  *string
	CreatorStar      int
	TasteStar        int
}

type Diaries struct {
	Items []Diary
}

type Fravor struct {
	Base
	Name       string
	Brand      FravorBrand
	CreateArea string
}

type Bottle struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	BottleBrand uuid.UUID `json:"brand_id"`
}

type Bowl struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	BowlBrand uuid.UUID `json:"brand_id"`
}

type HeatManagement struct {
	Base
	Name               string
	HeatManagemntBrand HeatManagementBrand
}

type Chacoal struct {
	Base
	Name         string
	ChacoalBrand ChacoalBrand
}

type FravorBrand struct {
	Base
	Name string
}

type BottleBrand struct {
	Base
	Name string
}

type BowlBrand struct {
	Base
	Name string
}

type HeatManagementBrand struct {
	Base
	Name string
}

type ChacoalBrand struct {
	Base
	Name string
}
