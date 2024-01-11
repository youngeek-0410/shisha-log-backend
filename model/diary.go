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
	Fravor           []Fravor       `gorm:"many2many:diary_fravors;embedded;embeddedPrefix:fravor_"`
	Bottle           Bottle         `gorm:"many2many:diary_bottles"`
	Bowl             Bowl           `gorm:"many2many:diary_bowls"`
	HeatManegement   HeatManegement `gorm:"many2many:diary_heat_manegements"`
	Chacoal          Chacoal        `gorm:"many2many:diary_chacoals"`
	Temperature      *int
	Humidity         *int
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
	Base
	Name        string
	BottleBrand BottleBrand
}

type Bowl struct {
	Base
	Name      string
	BowlBrand BowlBrand
}

type HeatManegement struct {
	Base
	Name               string
	HeatManegemntBrand HeatManegementBrand
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

type HeatManegementBrand struct {
	Base
	Name string
}

type ChacoalBrand struct {
	Base
	Name string
}
