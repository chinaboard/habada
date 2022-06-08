package model

import "gorm.io/gorm"

type HabadaItem struct {
	gorm.Model
	TinyUrl string `json:"tiny_url"  gorm:"unique;not null"`
	LongUrl string `json:"long_url"  gorm:"not null"`
}
