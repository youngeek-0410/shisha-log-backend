package bowl

import (
	"shisha-log-backend/lib"
	"shisha-log-backend/model"
)

type Bowls struct {
	Items []model.Bowl
}

func New() *Bowls {
	return &Bowls{}
}

// func (r *Bottles) Add(d model.Bottle) {
// 	r.Items = append(r.Items, d)
// 	db := lib.GetDBConn().DB
// 	if err := db.Create(d).Error; err != nil {
// 		fmt.Println("err!")
// 	}
// }

func (r *Bowls) GetAll() []model.Bowl {
	db := lib.GetDBConn().DB
	var bowls []model.Bowl
	if err := db.Find(&bowls).Error; err != nil {
		return nil
	}
	return bowls
}
