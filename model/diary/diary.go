package diary

import (
	"encoding/json"
	"shisha-log-backend/lib"
	"shisha-log-backend/model/flavor"
	"time"
)

type UserDiary struct {
	ID                string               `gorm:"column:id" json:"id"`
	DiaryFlavors      []flavor.DiaryFlavor `gorm:"foreignKey:ID" json:"diary_flavor_list"`
	CreateDate        time.Time            `gorm:"column:create_date" json:"create_date"`
	CreatorEvaluation float64              `gorm:"column:creator_evaluation" json:"creator_evaluation"`
	TasteEvaluation   float64              `gorm:"column:taste_evaluation" json:"taste_evaluation"`
}

type UserDiaries struct {
	Items []UserDiary `json:"user_diary_list"`
}

type DiaryRequest struct {
	UserID          string          `json:"user_id"`
	Equipments      DiaryEquipments `json:"equipments"`
	DiaryFlavorList []FlavorAmount  `json:"diary_flavor_list"`
	Image           string          `json:"image"`
	ServeText       string          `json:"serve_text"`
	SuckingText     string          `json:"sucking_text"`
	Review          DiaryReview     `json:"review"`
	CreateDate      string          `json:"create_date"`
}

type DiaryEquipments struct {
	BowlID           string  `json:"bowl_id"`
	BottleID         string  `json:"bottle_id"`
	HeatManagementID string  `json:"heat_management_id"`
	CharcoalID       string  `json:"charcoal_id"`
	Climate          Climate `json:"climate"`
}

type Climate struct {
	Temperature float64 `json:"temperature"`
	Humidity    float64 `json:"humidity"`
}

type FlavorAmount struct {
	ID     string  `json:"id"`
	Amount float64 `json:"amount"`
}

type DiaryReview struct {
	CreatorEvaluation float64 `json:"creator_evaluation"`
	TasteEvaluation   float64 `json:"taste_evaluation"`
	CreatorGoodPoints string  `json:"creator_good_points"`
	CreatorBadPoints  string  `json:"creator_bad_points"`
	TasteComments     string  `json:"taste_comments"`
}

type Diary struct {
	ID                string     `gorm:"column:id"`
	DiaryEquipmentsID string     `gorm:"column:diary_equipments_id"`
	ServeText         *string    `gorm:"column:serve_text"`
	SuckingText       *string    `gorm:"column:sucking_text"`
	Temperature       *float64   `gorm:"column:temperature"`
	Humidity          *float64   `gorm:"column:humidity"`
	CreatorEvaluation float64    `gorm:"column:creator_evaluation"`
	TasteEvaluation   float64    `gorm:"column:taste_evaluation"`
	CreatorGoodPoints *string    `gorm:"column:creator_good_points"`
	CreatorBadPoints  *string    `gorm:"column:creator_bad_points"`
	TasteComments     *string    `gorm:"column:taste_comments"`
	CreateDate        string     `gorm:"column:create_date"`
	CreatedAt         *time.Time `gorm:"column:created_at"`
	UpdatedAt         *time.Time `gorm:"column:updated_at"`
}

type Diaries struct {
	Items []Diary
}

func NewUserDiaries() *UserDiaries {
	return &UserDiaries{}
}

func (ud UserDiary) MarshalJSON() ([]byte, error) {
	type Alias UserDiary
	return json.Marshal(&struct {
		*Alias
		CreateDate string `json:"create_date"`
	}{
		Alias:      (*Alias)(&ud),
		CreateDate: ud.CreateDate.Format("2006-01-02"),
	})
}

func (r *UserDiaries) UserDiaries(user_id string) ([]UserDiary, error) {
	db := lib.GetDBConn().DB
	var userDiaries []UserDiary

	if err := db.Table("diaries").Select("diaries.id, diaries.create_date, diaries.creator_evaluation, diaries.taste_evaluation").Joins("inner join user_diaries on user_diaries.diary_id = diaries.id").Where("user_diaries.user_id = ?", user_id).Find(&userDiaries).Error; err != nil {
		return nil, err
	}

	for i := range userDiaries {
		var diaryFlavors flavor.DiaryFlavors

		flavors, err := diaryFlavors.DiaryFlavors(userDiaries[i].ID)
		if err != nil {
			return nil, err
		}

		userDiaries[i].DiaryFlavors = flavors
	}

	return userDiaries, nil
}

func (r *Diaries) Add(d Diary) error {
	r.Items = append(r.Items, d)
	db := lib.GetDBConn().DB
	if err := db.Create(&d).Error; err != nil {
		return err
	}
	return nil
}
