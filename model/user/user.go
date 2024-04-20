package user

import (
	"shisha-log-backend/lib"
	"time"
)

type UserDiary struct {
	ID        string
	UserID    string
	DiaryID   string
	CreatedAt *time.Time
}

type UserDiaries struct {
	Items []UserDiary
}

func (r *UserDiaries) Add(d UserDiary) error {
	r.Items = append(r.Items, d)
	db := lib.GetDBConn().DB
	if err := db.Create(&d).Error; err != nil {
		return err
	}
	return nil
}
