package image

import (
	"shisha-log-backend/lib"
	"time"
)

type DiaryImage struct {
	ID        string     `gorm:"column:id"`
	Path      string     `gorm:"column:path"`
	CreatedAt *time.Time `gorm:"column:created_at"`
	UpdatedAt *time.Time `gorm:"column:updated_at"`
}

type DiaryImages struct {
	Items []DiaryImage
}

func NewDiaryImages() *DiaryImages {
	return &DiaryImages{}
}

func (r *DiaryImages) Add(d DiaryImage) error {
	r.Items = append(r.Items, d)
	db := lib.GetDBConn().DB
	if err := db.Create(&d).Error; err != nil {
		return err
	}
	return nil
}
