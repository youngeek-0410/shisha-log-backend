package diary

import (
	"fmt"
	"shisha-log-backend/lib"
	"shisha-log-backend/model"
	"shisha-log-backend/model/bottle"
	"shisha-log-backend/model/bowl"
	"shisha-log-backend/model/charcoal"
	"shisha-log-backend/model/fravor"
	"shisha-log-backend/model/heatmanagement"
)

type Diary struct {
	Fravor           fravor.Fravors
	Bottle           bottle.Bottle
	Bowl             bowl.Bowl
	HeatManagement   heatmanagement.HeatManagement
	Chacoal          charcoal.Charcoal
	Temperature      *float64
	Humidity         *float64
	CreatorGoodPoint *string
	CreatorBadPoint  *string
	CreatorStar      int
	TasteStar        int
}

type userDiary struct {
}

type FlavorMix struct {
}

type Diaries struct {
	Items []model.Diary
}

func New() *Diaries {
	return &Diaries{}
}

func (r *Diaries) Add(d model.Diary) {
	r.Items = append(r.Items, d)
	db := lib.GetDBConn().DB
	if err := db.Create(d).Error; err != nil {
		fmt.Println("err!")
	}
}

func (r *Diaries) GetAll() []model.Diary {
	db := lib.GetDBConn().DB
	var diaries []model.Diary
	if err := db.Find(&diaries).Error; err != nil {
		return nil
	}
	return diaries
}
