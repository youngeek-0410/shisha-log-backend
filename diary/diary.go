package diary

import (
	"fmt"

	"shisha-log-backend/lib"
)

type Diary struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
}

type Diaries struct {
	Items []Diary
}

func New() *Diaries {
	return &Diaries{}
}

func (r *Diaries) Add(d Diary) {
	r.Items = append(r.Items, d)
	db := lib.GetDBConn().DB
	if err := db.Create(d).Error; err != nil {
		fmt.Println("err!")
	}
}

func (r *Diaries) GetAll() []Diary {
	db := lib.GetDBConn().DB
	var diaries []Diary
	if err := db.Find(&diaries).Error; err != nil {
		return nil
	}
	return diaries
}
