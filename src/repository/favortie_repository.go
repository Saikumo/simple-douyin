package repository

import (
	"gorm.io/gorm"
	"sync"
)

var (
	favoriteRepository     *FavoriteRepository
	favoriteRepositoryOnce sync.Once
)

type FavoriteRepository struct {
}

//单例获取点赞DAO
func GetFavoriteRepository() *FavoriteRepository {
	favoriteRepositoryOnce.Do(func() {
		favoriteRepository = new(FavoriteRepository)
	})
	return favoriteRepository
}

func (repository *FavoriteRepository) UserFavorite(userId, videoId uint) error {
	return DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Exec("UPDATE videos SET favorite_count=favorite_count+1 WHERE id = ?", videoId).Error; err != nil {
			return err
		}
		if err := tx.Exec("INSERT INTO `user_favorite_videos` (`user_id`,`video_id`) VALUES (?,?)", userId, videoId).Error; err != nil {
			return err
		}
		return nil
	})
}

func (repository *FavoriteRepository) UserUndoFavorite(userId, videoId uint) error {
	return DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Exec("UPDATE videos SET favorite_count=favorite_count-1 WHERE id = ? AND favorite_count>0", videoId).Error; err != nil {
			return err
		}
		if err := tx.Exec("DELETE FROM `user_favorite_videos`  WHERE `user_id` = ? AND `video_id` = ?", userId, videoId).Error; err != nil {
			return err
		}
		return nil
	})
}
