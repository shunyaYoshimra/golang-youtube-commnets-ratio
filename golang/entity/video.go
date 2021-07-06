package entity

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Video struct {
	ID        int       `gorm:"primaryKey" json:"id"`
	Title     string    `json:"title"`
	Ration    float64   `json:"ration"`
	VideoID   string    `json:"video_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func MigrateVideo(conn *gorm.DB) {
	conn.AutoMigrate(&Video{})
}
