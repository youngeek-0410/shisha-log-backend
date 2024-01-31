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
	Flavor []Flavor
	// Bottle bottle.Bottle
	// Bowl             bowl.Bowl
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

type Flavor struct {
	Base
	Name       string
	Brand      FlavorBrand
	CreateArea string
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

type FlavorBrand struct {
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
