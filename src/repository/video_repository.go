package repository

import (
	"saikumo.org/simple-douyin/src/entity"
	"sync"
	"time"
)

var (
	videoRepository     *VideoRepository
	VideoRepositoryOnce sync.Once
)

type VideoRepository struct {
}

func GetVideoRepository() *VideoRepository {
	VideoRepositoryOnce.Do(func() {
		videoRepository = new(VideoRepository)
	})
	return videoRepository
}

func (repository *VideoRepository) FindVideoByLimitAndLatestTime(limit int, latestTime time.Time) ([]*entity.Video, error) {
	var videoList []*entity.Video
	err := DB.Where("created_at<?", latestTime).
		Order("created_at DESC").
		Limit(limit).
		Find(&videoList).Error
	return videoList, err
}

func (repository *VideoRepository) FindVideoByUserId(userId uint) ([]*entity.Video, error) {
	var videoList []*entity.Video
	err := DB.Where("user_id=?", userId).
		Order("created_at DESC").
		Find(&videoList).Error
	return videoList, err
}
