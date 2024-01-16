package diary

import (
	"fmt"

	"shisha-log-backend/lib"
	"shisha-log-backend/model"
)

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
