package bottle

import (
	"shisha-log-backend/lib"
	"shisha-log-backend/model"
)

type Bottles struct {
	Items []model.Bottle
}

func New() *Bottles {
	return &Bottles{}
}

// func (r *Bottles) Add(d model.Bottle) {
// 	r.Items = append(r.Items, d)
// 	db := lib.GetDBConn().DB
// 	if err := db.Create(d).Error; err != nil {
// 		fmt.Println("err!")
// 	}
// }

func (r *Bottles) GetAll() []model.Bottle {
	db := lib.GetDBConn().DB
	var bottles []model.Bottle
	if err := db.Find(&bottles).Error; err != nil {
		return nil
	}
	return bottles
}
