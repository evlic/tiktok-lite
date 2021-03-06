package dao

import (
	"gorm.io/gorm"
	"tiktok-lite/global"
	"tiktok-lite/model"
	"tiktok-lite/util"
	"time"
)

// PublishActionDao 视频投稿，将视频信息持久化到数据库中
func PublishActionDao(user model.User, playUrl string, coverUrl string, title string) (int64, error) {
	now := time.Now()
	video := &model.Video{
		VideoId:       util.UniqueID(),
		Model:         gorm.Model{CreatedAt: now, UpdatedAt: now},
		UserId:        user.Id,
		PlayUrl:       playUrl,
		CoverUrl:      coverUrl,
		FavoriteCount: 0,
		CommentCount:  0,
		Title:         title,
	}

	if err := global.DB.Create(video).Error; err != nil {
		return 0, err
	}

	return video.VideoId, nil
}

// PublishList 查询用户发布视频列表
func PublishList(userId int64) ([]model.Video, error) {
	db := global.DB
	var videos []model.Video
	err := db.Where("user_id = ?", userId).Find(&videos).Error
	if err != nil {
		return videos, err
	}
	videoProcess(videos)
	return videos, nil
}

// PublishIdList 查询用户发布的视频 id 列表
func PublishIdList(userId int64) ([]int64, error) {
	db := global.DB
	var videosId []int64
	err := db.Table("video").Select("video_id").Where("user_id = ?", userId).Find(&videosId).Error
	if err != nil {
		return videosId, err
	}
	return videosId, nil
}
