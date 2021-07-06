package repositories

import (
	"github.com/jinzhu/gorm"
	"github.com/shunyaYoshimura/YouTube_Comments/database"
	"github.com/shunyaYoshimura/YouTube_Comments/entity"
)

type VideoRepository struct {
	Conn *gorm.DB
}

func NewVideoRepository() VideoRepository {
	return VideoRepository{Conn: database.GetDB().Table("videos")}
}

func (vr *VideoRepository) RetrieveVideos() (videos []entity.Video) {
	vr.Conn.Limit(10).Order("ration desc").Find(&videos)
	return
}

func (vr *VideoRepository) Create(video *entity.Video) (err error) {
	err = vr.Conn.Create(video).Error
	return
}
